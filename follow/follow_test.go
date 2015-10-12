package follow

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
	defer tx.Rollback()
	defer db.Close()
	User_1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	assert.Equal(t, "You have successfully followed red", User_1.Follow("red"), "People should be able to follow their person of choice")
}

func TestUser_Follow_1(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer tx.Rollback()
	defer db.Close()
	User_1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	assert.Equal(t, "You cannot follow an user who does not exist", User_1.Follow("bed"), "People should be able to follow their person of choice")
}
