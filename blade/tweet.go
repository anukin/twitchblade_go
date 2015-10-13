package blade

import (
	_ "database/sql"
	_ "fmt"
)

func (u User) Tweet(msg string) string {
	//var tweet string
	u.Transaction.Exec("INSERT INTO tweets(username, tweet) VALUES($1, $2)", u.Name, msg)
	return "Successfullly tweeted"
}
