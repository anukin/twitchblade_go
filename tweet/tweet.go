package tweet

import (
	_ "database/sql"
	_ "fmt"
	"github.com/anukin/twitchblade/mylib"
)

type User mylib.User

func (u User) Tweet(msg string) string {
	//var tweet string
	u.Transaction.Exec("INSERT INTO tweets(username, tweet) VALUES($1, $2)", u.Name, msg)
	return "Successfullly tweeted"
}
