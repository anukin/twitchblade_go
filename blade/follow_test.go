package blade

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Follow(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	user1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	user2 := User{"red", "charizard", tx}
	assert.Equal(t, "You have successfully followed red", user1.Follow(user2), "People should be able to follow their person of choice")
	tx.Rollback()
	db.Close()
}

func TestUser_Follow_1(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer tx.Rollback()
	defer db.Close()
	user1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	user2 := User{"red", "charizard", tx}
	user1.Follow(user2)
	assert.Equal(t, "You have already followed this user", user1.Follow(user2), "People should be able to follow their person of choice once")
}

func TestUser_Follow_2(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer tx.Rollback()
	defer db.Close()
	user1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	user2 := User{"bed", "charizard", tx}
	assert.Equal(t, "You cannot follow an user who does not exist", user1.Follow(user2), "People should be able to follow existing users")
}
