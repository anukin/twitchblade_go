package login

import (
	"database/sql"
	_ "fmt"
	_ "github.com/anukin/twitchblade/mylib"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Login(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, _ := db.Begin()
	User_1 := User{Name: "anugrah", Password: "megamind", Transaction: tx}
	assert.Equal(t, "Welcome to Twitchblade", User_1.Login(), "User should be able to login to twitchblade")
}
