package blade

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStream(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer tx.Rollback()
	defer db.Close()
	user1 := User{Name: "red", Password: "charizard", Transaction: tx}
	msg := "I will be the very best"
	_, tweet_id := user1.Tweet(msg)
	tweets := make([]Tweetmodel, 0)
	assert.Equal(t, append(tweets, Tweetmodel{tweet_id, msg}), Stream("red", tx), "Stream should return all tweets by an user")
}
