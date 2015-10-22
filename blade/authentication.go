package blade

import (
	"database/sql"
	_ "fmt"
)

type AuthenticationService struct {
	User
	*sql.Tx
}

func (auth *AuthenticationService) Authenticate() string {
	var username, password string
	err := auth.QueryRow("SELECT name, password FROM users WHERE name=$1", auth.Name).Scan(&username, &password)
	if err == sql.ErrNoRows {
		auth.Exec("INSERT INTO users(name, password) VALUES($1, $2)", auth.Name, auth.Password)
		return "Successfully registered"
	} else {
		if auth.Name == username && auth.Password == password {
			return "Welcome to Twitchblade"
		} else {
			return "Your password is wrong, please try again!"
		}
	}
}
