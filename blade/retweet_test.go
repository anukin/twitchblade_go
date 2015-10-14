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
	User_1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	_, tweet_id := User_1.Tweet("lol")
	defer tx.Rollback()
	defer db.Close()
	User_2 := User{Name: "red", Password: "charizard", Transaction: tx}
	assert.Equal(t, "Successfully retweeted tweet by anugrah", User_2.Retweet(tweet_id), "User should be able to retweet")
}

func TestUser_Retweet_1(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	User_1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	_, tweet_id := User_1.Tweet("lol")
	defer tx.Rollback()
	defer db.Close()
	User_2 := User{Name: "red", Password: "charizard", Transaction: tx}
	User_2.Retweet(tweet_id)
	assert.Equal(t, "You have already retweeted this tweet", User_2.Retweet(tweet_id), "User should be able to retweet")
}
