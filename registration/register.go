package registration

import (
	"database/sql"
	"fmt"
)

type user struct {
	name        string
	password    string
	transaction *sql.Tx
}

func (u user) Register() string {
	var username string
	err := u.transaction.QueryRow("SELECT * from users WHERE name=$1", u.name).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		u.transaction.Query("INSERT INTO users(name, password) VALUES($1, $2)", u.name, u.password)
		return "Successfully registered"
	default:
		fmt.Println(err)
		return "User exists with same name.Please try a new username"
	}
}
