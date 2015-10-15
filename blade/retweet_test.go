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
	message, _ := User_2.Retweet(tweet_id)
	assert.Equal(t, "Successfully retweeted tweet by anugrah", message, "User should be able to retweet")
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
	message, _ := User_2.Retweet(tweet_id)
	assert.Equal(t, "You have already retweeted this tweet", message, "User should be able to retweet the same tweet only once")
}

func TestUser_Retweet_2(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	User_1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	_, tweet_id := User_1.Tweet("lol")
	//defer tx.Rollback()
	tx.Commit()
	defer db.Close()
	User_2 := User{Name: "red", Password: "charizard", Transaction: tx}
	_, retweet_id := User_2.Retweet(tweet_id)
	tx.Commit()
	trax, _ := db.Begin()
	User_3 := User{Name: "lol", Password: "lol", Transaction: trax}
	message, _ := User_3.Retweet(retweet_id)
	tx.Rollback()
	assert.Equal(t, "Successfully retweeted tweet by anugrah", message, "User should be able to retweet those tweets of original tweet")
}
