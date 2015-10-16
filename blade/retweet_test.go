package blade

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestuserRetweet(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	user1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	_, tweet_id := user1.Tweet("lol")
	defer tx.Rollback()
	defer db.Close()
	user2 := User{Name: "red", Password: "charizard", Transaction: tx}
	message, _ := user2.Retweet(tweet_id)
	assert.Equal(t, "Successfully retweeted tweet by anugrah", message, "User should be able to retweet")
}

func TestuserRetweet_1(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	user1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	_, tweet_id := user1.Tweet("lol")
	defer tx.Rollback()
	defer db.Close()
	user2 := User{Name: "red", Password: "charizard", Transaction: tx}
	user2.Retweet(tweet_id)
	message, _ := user2.Retweet(tweet_id)
	assert.Equal(t, "You have already retweeted this tweet", message, "User should be able to retweet the same tweet only once")
}

func TestuserRetweet_2(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	user1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	_, tweet_id := user1.Tweet("lol")
	user2 := User{Name: "red", Password: "charizard", Transaction: tx}
	_, retweet_id := user2.Retweet(tweet_id)
	user3 := User{Name: "lol", Password: "lol", Transaction: tx}
	message, _ := user3.Retweet(retweet_id)
	assert.Equal(t, "Successfully retweeted tweet by anugrah", message, "User should be able to retweet those tweets of original tweet")
}
