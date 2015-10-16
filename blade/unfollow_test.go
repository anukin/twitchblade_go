package blade

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestuserUnfollow(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	user1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	user1.Follow("red")
	assert.Equal(t, "You have successfully unfollowed red", user1.Unfollow("red"), "People should be able to unfollow")
	tx.Rollback()
	db.Close()
}

func TestuserUnfollow_1(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	user1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	assert.Equal(t, "You do not follow this user", user1.Unfollow("red"), "People should be able to unfollow people whom they follow")
	tx.Rollback()
	db.Close()
}

// func TestuserUnfollow_2(t *testing.T) {
// 	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
// 	tx, err := db.Begin()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	user1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
// 	assert.Equal(t, "This user does not exist, you can unfollow only existing users.", user1.Unfollow("bed"), "People should be able to unfollow people who exist")
// 	tx.Rollback()
// 	db.Close()
// }
