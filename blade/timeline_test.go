package blade

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Timeline(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	msg := "Haha"
	user1 := User{Name: "lol", Password: "lol", Transaction: tx}
	_, tweetid := user1.Tweet(msg)
	user2 := User{Name: "red", Password: "charizard", Transaction: tx}
	user3 := User{"lol", "lol", tx}
	user2.Follow(user3)
	defer tx.Rollback()
	defer db.Close()
	tweets := make([]Tweetmodel, 0)
	assert.Equal(t, append(tweets, Tweetmodel{tweetid, msg}), user2.Timeline(), "Stream should return all tweets of those followed by user")
}
