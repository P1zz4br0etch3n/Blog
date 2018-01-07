/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestGenerateSession(t *testing.T) {
	GenerateSession("userName", "1234567890")
	users := GetOnlineUserNames()
	assert.Equal(t, users, []string{"userName"})
	os.RemoveAll(dataDir)
}