/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadUsers(t *testing.T) {
	LoadUsers()
	saveUsers()
	succ := VerifyUser("root", "toor")
	assert.True(t, succ == nil)
}

func TestAddUser(t *testing.T) {
	LoadUsers()
	AddUser("newUser", "möp")
	succ := VerifyUser("newUser", "möp")
	assert.True(t, succ == nil)
}

func TestVerifyUser(t *testing.T) {
	LoadUsers()
	succ := VerifyUser("idonotexist", "möp")
	assert.True(t, succ.Error() == "user does not exist")
}

func TestChangePassword(t *testing.T) {
	LoadUsers()
	ChangePassword("root", "toor", "test")
	succ := VerifyUser("root", "test")
	assert.True(t, succ == nil)
}