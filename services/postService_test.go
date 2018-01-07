/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
	"de/vorlesung/projekt/2416160-5836402/models"
	"time"
)

func TestLoadPosts(t *testing.T) {
	LoadPosts()
	post, err := GetMostRecentPost()
	assert.True(t, post.PostID == "0" && err == nil)
	os.RemoveAll(dataDir)
}

func TestNewPost(t *testing.T) {
	LoadPosts()
	NewPost(models.BlogPost{Content:"Post Content.", Author:"Auth"})
	post, err := GetMostRecentPost()
	assert.True(t, post.Author == "Auth" && post.Content == "Post Content.", post.PostID == "1" && err == nil)
	os.RemoveAll(dataDir)
}

func TestAppendCommentToPost(t *testing.T) {
	LoadPosts()
	NewPost(models.BlogPost{Content:"Post Content.", Author:"Auth"})
	AppendCommentToPost("1", &models.Comment{"Nick", time.Now(), "C0nt3nt"})
	post, err := GetMostRecentPost()
	assert.True(t, post.Comments[0].Nickname == "Nick" && post.Comments[0].Content == "C0nt3nt" && err == nil)
	os.RemoveAll(dataDir)
}