package blade

import (
	"database/sql"
	_ "fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Tweet(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, _ := db.Begin()
	msg := "Hello world!"
	defer tx.Rollback()
	defer db.Close()
	user1 := User{"anugrah", "megamind", tx}
	s, _ := user1.Tweet(msg)
	assert.Equal(t, "Successfullly tweeted", s, "User should be able to tweet")
}
