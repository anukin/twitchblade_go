package blade

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Retweet(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	user1 := User{"anugrah", "megamind", tx}
	_, tweetid := user1.Tweet("lol")
	defer tx.Rollback()
	defer db.Close()
	user2 := User{"red", "charizard", tx}
	message, _ := user2.Retweet(tweetid)
	assert.Equal(t, "Successfully retweeted tweet by anugrah", message, "User should be able to retweet")
}

func TestUser_Retweet_1(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	user1 := User{"anugrah", "megamind", tx}
	_, tweetid := user1.Tweet("lol")
	user2 := User{"red", "charizard", tx}
	user2.Retweet(tweetid)
	message, _ := user2.Retweet(tweetid)
	tx.Rollback()
	db.Close()
	assert.Equal(t, "You have already retweeted this tweet", message, "User should be able to retweet the same tweet only once")
}

func TestUser_Retweet_2(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	user1 := User{"anugrah", "megamind", tx}
	_, tweetid := user1.Tweet("lol")
	user2 := User{"red", "charizard", tx}
	_, retweetid := user2.Retweet(tweetid)
	user3 := User{"lol", "lol", tx}
	message, _ := user3.Retweet(retweetid)
	tx.Rollback()
	db.Close()
	assert.Equal(t, "Successfully retweeted tweet by anugrah", message, "User should be able to retweet those tweets of original tweet")
}
