package blade

import (
	"database/sql"
	_ "fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Login(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, _ := db.Begin()
	user1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	assert.Equal(t, "Welcome to Twitchblade", user1.Login(), "User should be able to login to twitchblade")
}

func TestUser_Login_1(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, _ := db.Begin()
	user2 := User{Name: "bobo", Password: "superboy", Transaction: tx}
	assert.Equal(t, "There is no user with that name, please try again or try registering!", user2.Login(), "Only Valid user should be able to use twitchblade")
}

func TestUser_Login_2(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, _ := db.Begin()
	user3 := User{Name: "anugrah", Password: "superboy", Transaction: tx}
	assert.Equal(t, "Your password is wrong, please try again!", user3.Login(), "Users should be able to login with valid credentials")
}
