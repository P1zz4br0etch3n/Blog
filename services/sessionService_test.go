/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import "testing"

func TestGenerateSession(t *testing.T) {
	GenerateSession("user", "1234567890")
}