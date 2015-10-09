package login

import (
	_ "database/sql"
	_ "fmt"
	"github.com/anukin/twitchblade/mylib"
)

type User mylib.User

func (u User) Login() string {
	return "Welcome to Twitchblade"
}
