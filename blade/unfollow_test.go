package blade

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Unfollow(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	User_1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	User_1.Follow("red")
	assert.Equal(t, "You have successfully unfollowed red", User_1.Unfollow("red"), "People should be able to unfollow")
	tx.Rollback()
	db.Close()
}

func TestUser_Unfollow_1(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	User_1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	assert.Equal(t, "You do not follow this user", User_1.Unfollow("red"), "People should be able to unfollow people whom they follow")
	tx.Rollback()
	db.Close()
}
