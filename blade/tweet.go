package blade

import (
	_ "database/sql"
	_ "fmt"
)

type Tweetmodel struct {
	id      int
	message string
}

func (u User) Tweet(msg string) (string, int) {
	var id int
	u.Transaction.QueryRow("INSERT INTO tweets(username, tweet) VALUES($1, $2) returning id", u.Name, msg).Scan(&id)
	return "Successfullly tweeted", id
}
