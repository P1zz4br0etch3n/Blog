/*
    Matrikelnummern: 5836402, 2416160
*/

package models

import (
	"time"
)

type User struct {
	Username string
	Password string
}

type Comment struct {
	Text   string
	Time   time.Time
	Author string
}

type Blog struct {
	Id       string
	Title    string
	Text     string
	Time     time.Time
	Author   *User
	Comments []*Comment
}

type Settings struct {
	PortNumber     string
	SessionTimeout uint

	PostDirectory string
	PostSuffix    string
	//NumberOfPosts  uint

	KeyDirectory string
	KeyFile      string
	CertFile     string

	TemplateDirectory string
	TemplateSuffix    string
}
