/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"de/vorlesung/projekt/2416160-5836402/models"
	"errors"
)

func AuthenticateUser(uname, pw string) error {
	var u models.User
	e := ReadJsonFile(uname, UserDir, &u)
	if e != nil {
		return e
	}
	if uname == u.Username && pw == u.Password {
		return nil
	}
	return errors.New("Password is not correct")
}
