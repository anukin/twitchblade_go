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
		return fmt.Sprintf("You have successfully followed %v", name)
	}
}
