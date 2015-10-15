package blade

import (
	_ "database/sql"
	"fmt"
)

func (u *User) Timeline() []Tweetmodel {
	rows, err := u.Transaction.Query("select tweets.id, tweets.tweet from tweets INNER JOIN follow ON (tweets.username = follow.following) and follow.username=$1", u.Name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	tweets := make([]Tweetmodel, 0)
	for rows.Next() {
		var id int
		var msg string
		rows.Scan(&id, &msg)
		fmt.Println(id, msg)
		tweet := Tweetmodel{id, msg}
		tweets = append(tweets, tweet)
	}
	return tweets
}
