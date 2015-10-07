package registration

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Register(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer tx.Rollback()
	defer db.Close()
	User_1 := user{name: "bobo", password: "lol", transaction: tx}
	assert.Equal(t, "Successfully registered", User_1.Register(), "User should be able to register successfully")
	User_2 := user{name: "bobo", password: "lol", transaction: tx}
	assert.Equal(t, "User exists with same name.Please try a new username", User_2.Register(), "User name should be unique")
}
