/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"encoding/hex"
	"crypto/rand"
	"log"
	"errors"
	"net/http"
	"de/vorlesung/projekt/2416160-5836402/global"
	"time"
	"de/vorlesung/projekt/2416160-5836402/models"
	"sync"
)

var sessionMap = make(map[string]*models.Session)
var lock sync.Mutex

func GetOnlineUserNames() []string {
	var names []string
	for _, value := range sessionMap {
		names = append(names, value.UserName)
	}
	return names
}

func getSessionTimeout() time.Duration {
	return time.Duration(global.Settings.SessionTimeout) * time.Minute
}

func DestroySession(r *http.Request) {
	cookie, err := GetSessionCookie(r)
	if err == nil {
		lock.Lock()
		defer lock.Unlock()

		delete(sessionMap, cookie.Value)
	}
}

func GenerateSession(username, sessionId string) {
	lock.Lock()
	defer lock.Unlock()

	timer := time.AfterFunc(getSessionTimeout(), func() {
		lock.Lock()
		defer lock.Unlock()

		delete(sessionMap, sessionId)
	})

	session := models.Session{
		UserName: username,
		Expires:  time.Now().Add(getSessionTimeout()),
		Timer:    timer,
	}

	sessionMap[sessionId] = &session
}

func GetSessionCookie(r *http.Request) (*http.Cookie, error) {
	return r.Cookie("session")
}

func CheckSession(r *http.Request) (*models.Session, error) {
	cookie, err := GetSessionCookie(r)
	if err != nil {
		return nil, err
	}
	session := sessionMap[cookie.Value]
	if session != nil {
		session.Timer.Reset(getSessionTimeout())
		return session, nil
	}
	return nil, errors.New("Session invalid.")
}

func GenerateCookie() (*http.Cookie, error) {
	sessionId, e := generateSessionId()
	if e != nil {
		return nil, e
	}

	cookie := http.Cookie{
		Name:    "session",
		Value:   sessionId,
		MaxAge:  0,
		Expires: time.Now().Add(365 * 24 * time.Hour), // expires in one year
		Path:    "/",
		Secure:  true,
	}

	return &cookie, nil
}

func generateSessionId() (string, error) {
	data := make([]byte, 128)
	_, err := rand.Read(data)

	if err != nil {
		log.Println("Could not create session id.")
		return "", errors.New("Could not create session id.")
	}

	return hex.EncodeToString(data), nil
}
