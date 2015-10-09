package mylib

import "database/sql"

type User struct {
	Name        string
	Password    string
	Transaction *sql.Tx
}
