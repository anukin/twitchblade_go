package blade

import (
	"database/sql"
	"fmt"
)

type User struct {
	Name     string
	Password string
	*sql.Tx
}

func Stream(username string, transaction *sql.Tx) []Tweetmodel {
	rows, err := transaction.Query("Select id, tweet from tweets where username=$1", username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	tweets := make([]Tweetmodel, 0)
	for rows.Next() {
		var id int
		var msg string
		rows.Scan(&id, &msg)
		tweet := Tweetmodel{id, msg}
		tweets = append(tweets, tweet)
	}
	return tweets
}

func (u *User) Follow(usertofollow User) string {
	name := usertofollow.Name
	var username string
	err := u.QueryRow("SELECT name from users where name=$1", name).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		return "You cannot follow an user who does not exist"
	default:
		if u.alreadyfollowing(name) {
			return "You have already followed this user"
		} else {
			u.Exec("INSERT INTO follow(username, following) VALUES($1, $2)", u.Name, name)
			return fmt.Sprintf("You have successfully followed %v", name)
		}
	}
}

func (u *User) alreadyfollowing(usertofollow string) bool {
	res, _ := u.Query("SELECT * from follow where username=$1 and following=$2", u.Name, usertofollow)
	return (res.Next() == true)
}

func (u User) Login() string {
	var username, password string
	err := u.QueryRow("SELECT name, password FROM users WHERE name=$1", u.Name).Scan(&username, &password)
	if err == sql.ErrNoRows {
		return "There is no user with that name, please try again or try registering!"
	}
	if u.Password != password {
		return "Your password is wrong, please try again!"
	}
	return "Welcome to Twitchblade"
}

func (u *User) Register() string {
	var username string
	err := u.QueryRow("SELECT name FROM users WHERE name=$1", u.Name).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		u.Exec("INSERT INTO users(name, password) VALUES($1, $2)", u.Name, u.Password)
		return "Successfully registered"
	default:
		return "User exists with same name.Please try a new username"
	}
}

func (u *User) alreadyretweeted(tweetid int) bool {
	var id int
	err := u.QueryRow("SELECT id from retweets where original_tweet_id = $1 and retweeted_by = $2", tweetid, u.Name).Scan(&id)
	return (err != sql.ErrNoRows)
}

func (u *User) iteratedretweet(tweetid int) (bool, int) {
	var id int
	err := u.QueryRow("SELECT original_tweet_id from retweets where retweet_tweet_id = $1", tweetid).Scan(&id)
	return (err != sql.ErrNoRows), id
}

func (u *User) Retweet(tweetid int) (string, int) {
	if u.alreadyretweeted(tweetid) {
		return "You have already retweeted this tweet", tweetid
	} else {
		flag, originalid := u.iteratedretweet(tweetid)
		if flag {
			return u.Retweet(originalid)
		} else {
			var msg, originaluser string
			var id int
			u.QueryRow("select username, tweet from tweets where id=$1", tweetid).Scan(&originaluser, &msg)
			_, retweetid := u.Tweet(msg)
			u.QueryRow("INSERT INTO retweets(original_tweet_id, retweeted_by, retweet_tweet_id) VALUES($1, $2, $3) returning id", tweetid, u.Name, retweetid).Scan(&id)
			return fmt.Sprintf("Successfully retweeted tweet by %s", originaluser), retweetid
		}
	}
}

func (u *User) Timeline() []Tweetmodel {
	rows, err := u.Query("select tweets.id, tweets.tweet from tweets INNER JOIN follow ON (tweets.username = follow.following) and follow.username=$1", u.Name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	tweets := make([]Tweetmodel, 0)
	for rows.Next() {
		var id int
		var msg string
		rows.Scan(&id, &msg)
		tweet := Tweetmodel{id, msg}
		tweets = append(tweets, tweet)
	}
	return tweets
}

func (u User) Tweet(msg string) (string, int) {
	var id int
	u.QueryRow("INSERT INTO tweets(username, tweet) VALUES($1, $2) returning id", u.Name, msg).Scan(&id)
	return "Successfullly tweeted", id
}

func (u *User) Unfollow(usertounfollow User) string {
	res, _ := u.Query("SELECT * from follow where username=$1 and following=$2", u.Name, usertounfollow.Name)
	if res.Next() != true {
		return "You do not follow this user"
	} else {
		u.Exec("DELETE FROM follow WHERE name=$1 and following=$2)", u.Name, usertounfollow.Name)
		return fmt.Sprintf("You have successfully unfollowed %v", usertounfollow.Name)
	}
}
