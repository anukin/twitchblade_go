package blade

import (
	"database/sql"
	"fmt"
)

func (u *User) Register() string {
	var username string
	err := u.Transaction.QueryRow("SELECT name FROM users WHERE name=$1", u.Name).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		u.Transaction.Exec("INSERT INTO users(name, password) VALUES($1, $2)", u.Name, u.Password)
		return "Successfully registered"
	default:
		fmt.Println(err)
		return "User exists with same name.Please try a new username"
	}
}
