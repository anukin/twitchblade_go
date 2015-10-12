package follow

import (
	"database/sql"
	"fmt"
	"github.com/anukin/twitchblade/mylib"
)

type User mylib.User

func (u *User) Follow(name string) string {
	var username string
	err := u.Transaction.QueryRow("SELECT name from users where name=$1", name).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		return "You cannot follow an user who does not exist"
	default:
		if u.alreadyfollowing(name) {
			return "You have already followed this user"
		} else {
			u.Transaction.Exec("INSERT INTO follow(username, following) VALUES($1, $2)", u.Name, name)
			return fmt.Sprintf("You have successfully followed %v", name)
		}
	}
}

func (u *User) alreadyfollowing(usertofollow string) bool {
	//var username, following string
	res, _ := u.Transaction.Query("SELECT * from follow where username=$1 and following=$2", u.Name, usertofollow)
	return (res.Next() == true)
}
