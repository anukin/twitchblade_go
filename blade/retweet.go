package blade

import (
	"database/sql"
	"fmt"
)

func (u *User) alreadyretweeted(tweetid int) bool {
	var id string
	fmt.Println(u.Name)
	err := u.Transaction.QueryRow("SELECT retweeted_by from retweets where original_tweet_id = $1 and retweeted_by = $2", tweetid, u.Name).Scan(&id)
	return (err != sql.ErrNoRows)
}

func (u *User) iteratedretweet(tweetid int) (bool, int) {
	var id int
	err := u.Transaction.QueryRow("SELECT original_tweet_id from retweets where retweet_tweet_id = $1", tweetid).Scan(&id)
	return (err != sql.ErrNoRows), id
}

func (u *User) Retweet(tweetid int) (string, int) {
	//return "Successfully retweeted tweet by anugrah"
	if u.alreadyretweeted(tweetid) {
		fmt.Println("you are here : ", u.Name)
		return "You have already retweeted this tweet", tweetid
	} else {
		fmt.Println("anybody in here", u.Name)
		flag, originalid := u.iteratedretweet(tweetid)
		if flag == false {
			fmt.Println("inside not iterated retweet : ", u.Name)
			var msg, originaluser string
			u.Transaction.QueryRow("Select username, tweet from tweets where id = $1", tweetid).Scan(&originaluser, &msg)
			_, retweetid := u.Tweet(msg)
			var retweet_tweet_id int
			u.Transaction.QueryRow("INSERT INTO retweets(original_tweet_id, retweeted_by, retweet_tweet_id)) VALUES($1,$2,$3) returning retweet_tweet_id", tweetid, u.Name, retweetid).Scan(&retweet_tweet_id)
			return fmt.Sprintf("Successfully retweeted tweet by %s", originaluser), retweet_tweet_id
		} else {
			fmt.Println("Haha iterated tweet : ", u.Name)
			return u.Retweet(originalid)
		}
	}
}
