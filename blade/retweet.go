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

func (u *User) iteratedretweet(tweetid int) (bool, int) {
	var id int
	err := u.Transaction.QueryRow("SELECT original_tweet_id from retweets where retweet_tweet_id = $1", tweetid).Scan(&id)
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
			u.Transaction.QueryRow("select username, tweet from tweets where id=$1", tweetid).Scan(&originaluser, &msg)
			_, retweetid := u.Tweet(msg)
			u.Transaction.QueryRow("INSERT INTO retweets(original_tweet_id, retweeted_by, retweet_tweet_id) VALUES($1, $2, $3) returning id", tweetid, u.Name, retweetid).Scan(&id)
			return fmt.Sprintf("Successfully retweeted tweet by %s", originaluser), retweetid
		}
	}
}
