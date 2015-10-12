package follow

import (
	_ "database/sql"
	"fmt"
	"github.com/anukin/twitchblade/mylib"
)

type User mylib.User

func (u *User) Follow(name string) string {
	return fmt.Sprintf("You have successfully followed %v", name)
}
