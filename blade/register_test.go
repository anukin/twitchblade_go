package blade

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
	user1 := User{"bobo", "lol", tx}
	assert.Equal(t, "Successfully registered", user1.Register(), "User should be able to register successfully")
}

func TestUser_Register_1(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer tx.Rollback()
	defer db.Close()
	user1 := User{"bobo", "lol", tx}
	user2 := User{"bobo", "lol", tx}
	user1.Register()
	assert.Equal(t, "User exists with same name.Please try a new username", user2.Register(), "User name should be unique")
}
