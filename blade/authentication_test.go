package blade

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthenticationService_Authenticate(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer tx.Rollback()
	defer db.Close()
	user1 := User{"bol", "bol", tx}
	auth := AuthenticationService{user1, tx}
	assert.Equal(t, "Successfully registered", auth.Authenticate(), "Authentication should be able to register users")
}

func TestAuthenticationService_Authenticate_1(t *testing.T) {
	db, _ := sql.Open("postgres", "user=CodeWalker dbname=twitchblade_test sslmode=disable")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer tx.Rollback()
	defer db.Close()
	user1 := User{"bol", "bol", tx}
	auth := AuthenticationService{user1, tx}
	auth.Authenticate()
	assert.Equal(t, "Welcome to Twitchblade", auth.Authenticate(), "Authentication should be able to register users")
}
