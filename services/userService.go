/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"de/vorlesung/projekt/2416160-5836402/models"
	"errors"
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"crypto/rand"
)

var users = []models.User{{
	UserName:     "root",
	PasswordHash: "2c00f030fdf7d6e61a5adc7d012e95fbe1f6e411d83a4a33612bf5f873bf72fa530ad4364939d1cd3e045929706530d74a8908fc3eac7a7c1e4da5fe476d054d",
	Salt:         "71cbc0451d1fca78172a816b98d085d709f9eb964b8b883df7063505d5e2b34f30f0162c0acaa3f0bf95450c405575c0eecd8d9e0c480b7b01ffca619e0b13eb",
}}
var usersLoaded = false

func LoadUsers() error {
	if usersLoaded {
		log.Println("Users already loaded.")
		return nil
	}

	err := ReadJsonFile(UserDir, "users", &users)
	if err != nil {
		log.Println("Could not read Users-File. Creating new file with default values...")
		e := saveUsers()
		if e != nil {
			return errors.New("Could not create default User-File.")
		}
	}
	usersLoaded = true
	return nil
}

func saveUsers() error {
	err := WriteJsonFile(UserDir, "users", users)
	if err != nil {
		log.Println("Could not write to Users-File.")
		return err
	}
	return nil
}

func VerifyUser(userName, password string) error {
	for _, user := range users {
		if user.UserName == userName {
			salt := user.Salt
			genHash := calculateHash(password, salt)
			if genHash == user.PasswordHash {
				return nil
			} else {
				return errors.New("wrong password")
			}
		}
	}

	return errors.New("user does not exist")
}

func AddUser(userName, password string) error {
	if !usersLoaded {
		e := LoadUsers()
		if e != nil {
			log.Println("Could not add user. User-File not loaded.")
			return e
		}
	}

	for _, user := range users {
		if user.UserName == userName {
			return errors.New("username already taken")
		}
	}

	salt := generateSalt()
	hash := calculateHash(password, salt)

	newUser := models.User{
		UserName:     userName,
		PasswordHash: hash,
		Salt:         salt,
	}

	users = append(users, newUser)
	saveUsers()
	return nil
}

func ChangePassword(userName, oldPassword, newPassword string) error {
	if err := VerifyUser(userName, oldPassword); err == nil {
		for i, user := range users {
			if user.UserName == userName {
				salt := generateSalt()
				hash := calculateHash(newPassword, salt)

				users[i].PasswordHash = hash
				users[i].Salt = salt
			}
		}
		err = saveUsers()
		if err != nil {
			return errors.New("Couldn't change password. Save User-File failed.")
		}
		return nil
	} else {
		return errors.New("Couldn't change password. VerifyUser failed.")
	}
}

func calculateHash(password, salt string) string {
	saltedPasswordBytes := append([]byte(password + salt))

	hashFunc1 := md5.New()
	hashFunc1.Write(saltedPasswordBytes)

	hashFunc2 := sha512.New()
	hashFunc2.Write(append(hashFunc1.Sum(nil), saltedPasswordBytes...))

	return hex.EncodeToString(hashFunc2.Sum(nil))
}

func generateSalt() string {
	data := make([]byte, 32)
	_, err := rand.Read(data)

	if err != nil {
		fmt.Println("Could not read random data.")
	}

	hashFunc := sha512.New()
	hashFunc.Write(data)

	return hex.EncodeToString(hashFunc.Sum(nil))
}
