package login

import (
	"database/sql"
	_ "fmt"
	"github.com/anukin/twitchblade/mylib"
)

type User mylib.User

func (u User) Login() string {
	var username, password string
	err := u.Transaction.QueryRow("SELECT name, password FROM users WHERE name=$1", u.Name).Scan(&username, &password)
	if err == sql.ErrNoRows {
		return "There is no user with that name, please try again or try registering!"
	}
	if u.Password != password {
		return "Your password is wrong, please try again!"
	}
	return "Welcome to Twitchblade"
}
