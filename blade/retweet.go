package blade

import (
	"database/sql"
	"fmt"
)

func (u *User) alreadyretweeted(tweetid int) bool {
	var id int
	err := u.Transaction.QueryRow("SELECT id from retweets where original_tweet_id = $1 and retweeted_by = $2", tweetid, u.Name).Scan(&id)
	return (err != sql.ErrNoRows)
}

func (u *User) Retweet(tweetid int) string {
	//return "Successfully retweeted tweet by anugrah"
	if u.alreadyretweeted(tweetid) {
		return "You have already retweeted this tweet"
	} else {
		var msg, originaluser string
		u.Transaction.QueryRow("Select username, tweet from tweets where id = $1", tweetid).Scan(&originaluser, &msg)
		_, retweetid := u.Tweet(msg)
		u.Transaction.Exec("INSERT INTO retweets(original_tweet_id, retweeted_by, retweet_tweet_id)) VALUES($1,$2,$3)", tweetid, u.Name, retweetid)
		return fmt.Sprintf("Successfully retweeted tweet by %s", originaluser)
	}
}
