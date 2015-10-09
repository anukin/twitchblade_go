package login

import (
	"database/sql"
	_ "fmt"
	"github.com/anukin/twitchblade/mylib"
)

type User mylib.User

func (u User) Login() string {
	var username, password string
	err := u.Transaction.QueryRow("Select * from users where name=$1", u.Name).Scan(&username, &password)
	//fmt.Println(password)
	switch {
	case err == sql.ErrNoRows:
		return "There is no user with that name, please try again!"
	}
	return "Welcome to Twitchblade"
}
