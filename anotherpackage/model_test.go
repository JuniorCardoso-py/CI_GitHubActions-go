package user_adapter_test

import (
	"testing"

	user "github.com/golangci_test/anotherpackage"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	want := &user.User{
		Name: "Junior",
		Age:  26,
	}

	newUser := user.CreateUser("Junior", 26)

	assert.Equal(t, want, newUser.GetUser())
}
