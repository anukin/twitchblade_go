package userlogin

import (
    "database/sql"
    _"fmt"
    _ "github.com/lib/pq"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestUser_Login(t *testing.T) {
    User_1 = user(name: "anugrah", password: "megamind", transaction: tx)
    assert.Equal(t, "Welcome to Twitchblade", User_1.Login(), "User should be able to login to twitchblade")
}