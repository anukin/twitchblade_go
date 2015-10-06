package registration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistration_Register(t *testing.T) {
	User := user{name: "anukin", password: "lol"}
	assert.Equal(t, "Successfully registered", User.Register(), "User should be able to register successfully")
}
