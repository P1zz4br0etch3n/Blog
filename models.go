/*
    Matrikelnummern: 5836402, 2416160
*/

package main

import (
	"time"
)

type User struct {
	Username string
	Password string
}

type Comment struct {
	Text string
	Time time.Time
	Author string
}

type Blog struct {
	Id string
	Text string
	Time time.Time
	Author *User
	Comments []*Comment
}

