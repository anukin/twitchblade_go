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
	user1 := User{"red", "charizard", tx}
	msg := "I will be the very best"
	_, tweetid := user1.Tweet(msg)
	tweets := make([]Tweetmodel, 0)
	assert.Equal(t, append(tweets, Tweetmodel{tweetid, msg}), Stream("red", tx), "Stream should return all tweets by an user")
}
