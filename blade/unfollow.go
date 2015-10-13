package blade

import (
	_ "database/sql"
	"fmt"
)

//type User mylib.User

func (u *User) Unfollow(name string) string {
	//return "You have successfully unfollowed red"
	//var usertounfollow string
	res, _ := u.Transaction.Query("SELECT * from follow where username=$1 and following=$2", u.Name, name)
	//fmt.Println(err)
	if res.Next() != true {
		return "You do not follow this user"
	} else {
		u.Transaction.Exec("DELETE FROM follow WHERE name=$1 and following=$2)", u.Name, name)
		return fmt.Sprintf("You have successfully unfollowed %v", name)
	}
}
