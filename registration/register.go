package registration

import (
	"database/sql"
	"fmt"
	"github.com/anukin/twitchblade/mylib"
)

type User mylib.User

func (u *User) Register() string {
	var username string
	err := u.Transaction.QueryRow("SELECT * from users WHERE name=$1", u.Name).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		u.Transaction.Query("INSERT INTO users(name, password) VALUES($1, $2)", u.Name, u.Password)
		return "Successfully registered"
	default:
		fmt.Println(err)
		return "User exists with same name.Please try a new username"
	}
}
