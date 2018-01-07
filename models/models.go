/*
    Matrikelnummern: 5836402, 2416160
*/

package models

import (
	"time"
)

type User struct {
	UserName     string
	PasswordHash string
	Salt         string
}

type Comment struct {
	Nickname string
	Date     time.Time
	Content  string
}

type BlogPost struct {
	PostID   string
	Author   string
	Date     time.Time
	Content  string
	Comments []Comment
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

type Session struct {
	UserName string
	Expires  time.Time
	Timer    *time.Timer
}

type IndexPage struct {
	UserLoggedIn    bool
	ShowArchiveLink bool
	UserName        string
	Posts           []BlogPost
}
