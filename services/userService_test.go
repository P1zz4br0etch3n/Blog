/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestLoadUsers(t *testing.T) {
	LoadUsers()
	saveUsers()
	succ := VerifyUser("root", "toor")
	assert.True(t, succ == nil)
	os.RemoveAll(dataDir)
}

func TestAddUser(t *testing.T) {
	LoadUsers()
	AddUser("newUser", "möp")
	succ := VerifyUser("newUser", "möp")
	assert.True(t, succ == nil)
	os.RemoveAll(dataDir)
}

func TestVerifyUser(t *testing.T) {
	LoadUsers()
	succ := VerifyUser("idonotexist", "möp")
	assert.True(t, succ.Error() == "user does not exist")
	os.RemoveAll(dataDir)
}

func TestVerifyUser2(t *testing.T) {
	LoadUsers()
	succ := VerifyUser("root", "möp")
	assert.True(t, succ.Error() == "wrong password")
	os.RemoveAll(dataDir)
}

func TestChangePassword(t *testing.T) {
	LoadUsers()
	ChangePassword("root", "toor", "test")
	succ := VerifyUser("root", "test")
	assert.True(t, succ == nil)
	os.RemoveAll(dataDir)
}