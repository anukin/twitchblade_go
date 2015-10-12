package tweet

import (
	"database/sql"
	_ "fmt"
	_ "github.com/anukin/twitchblade/mylib"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Tweet(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, _ := db.Begin()
	msg := "Hello world!"
	defer tx.Rollback()
	//defer tx.Commit()
	defer db.Close()
	User_1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	assert.Equal(t, "Successfullly tweeted", User_1.Tweet(msg), "User should be able to tweet")
}
