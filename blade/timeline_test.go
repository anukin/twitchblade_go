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
	User_1 := User{Name: "lol", Password: "lol", Transaction: tx}
	_, tweet_id := User_1.Tweet(msg)
	//tx.Commit()
	// trax, _ := db.Begin()
	User_2 := User{Name: "red", Password: "charizard", Transaction: tx}
	User_2.Follow("lol")
	defer tx.Rollback()
	defer db.Close()
	tweets := make([]Tweetmodel, 0)
	assert.Equal(t, append(tweets, Tweetmodel{tweet_id, msg}), User_2.Timeline(), "Stream should return all tweets of those followed by user")
}
