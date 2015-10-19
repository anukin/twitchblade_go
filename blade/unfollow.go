package blade

import (
	_ "database/sql"
	"fmt"
)

func (u *User) Unfollow(usertounfollow User) string {
	res, _ := u.Transaction.Query("SELECT * from follow where username=$1 and following=$2", u.Name, usertounfollow.Name)
	if res.Next() != true {
		return "You do not follow this user"
	} else {
		u.Transaction.Exec("DELETE FROM follow WHERE name=$1 and following=$2)", u.Name, usertounfollow.Name)
		return fmt.Sprintf("You have successfully unfollowed %v", usertounfollow.Name)
	}
}
