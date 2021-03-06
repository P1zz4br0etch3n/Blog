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

func PrepareTests() {
	SetPostManagerSettings(".json")
	posts = []models.BlogPost{}
	lastPostID = 0
	postsLoaded = false
}

func TestLoadPosts(t *testing.T) {
	PrepareTests()

	LoadPosts()
	NewPost(models.BlogPost{Content:"Post Content.", Author:"Auth"})
	post, err := GetMostRecentPost()
	assert.True(t, post.PostID == "1" && err == nil)
	os.RemoveAll(dataDir)
}

func TestNewPost(t *testing.T) {
	PrepareTests()

	LoadPosts()
	NewPost(models.BlogPost{Content:"Post Content.", Author:"Auth"})
	post, err := GetMostRecentPost()
	assert.True(t, post.Author == "Auth" && post.Content == "Post Content." && post.PostID == "1" && err == nil)
	os.RemoveAll(dataDir)
}

func TestAppendCommentToPost(t *testing.T) {
	PrepareTests()

	LoadPosts()
	NewPost(models.BlogPost{Content:"Post Content.", Author:"Auth"})
	AppendCommentToPost("1", &models.Comment{"Nick", time.Now(), "C0nt3nt"})
	post, err := GetMostRecentPost()
	assert.True(t, post.Comments[0].Nickname == "Nick" && post.Comments[0].Content == "C0nt3nt" && err == nil)
	os.RemoveAll(dataDir)
}

func TestDeletePost(t *testing.T) {
	PrepareTests()

	LoadPosts()
	NewPost(models.BlogPost{Content:"Post Content.", Author:"Auth"})
	DeletePost("1")
	post, err := GetMostRecentPost()
	assert.True(t, post.PostID == "0" && err == nil)
	os.RemoveAll(dataDir)
}

func TestChangePost(t *testing.T) {
	PrepareTests()

	LoadPosts()
	NewPost(models.BlogPost{Content:"Post Content.", Author:"Auth"})
	ChangePost("1", "Im No Longer 'Post Content.'")
	post, err := GetMostRecentPost()
	assert.True(t, post.Author == "Auth" && post.PostID == "1" && err == nil)
	assert.True(t, post.Content == "Im No Longer 'Post Content.'")
	os.RemoveAll(dataDir)
}

func TestGetAllPosts(t *testing.T) {
	PrepareTests()

	LoadPosts()
	NewPost(models.BlogPost{Content:"Post Content. (1)", Author:"Auth"})
	NewPost(models.BlogPost{Content:"Post Content. (2)", Author:"Auth"})
	NewPost(models.BlogPost{Content:"Post Content. (3)", Author:"Auth"})
	post, err := GetAllPosts()
	assert.True(t, len(post) == 3 && err == nil)
	os.RemoveAll(dataDir)
}

func TestGetAllPostsFromUser(t *testing.T) {
	PrepareTests()

	LoadPosts()
	NewPost(models.BlogPost{Content:"Post Content. (1)", Author:"Auth"})
	NewPost(models.BlogPost{Content:"Post Content. (2)", Author:"AuthorAlternative"})
	NewPost(models.BlogPost{Content:"Post Content. (3)", Author:"Auth"})
	post, err := GetAllPostsFromUser("Auth")
	assert.True(t, len(post) == 2 && err == nil)
	os.RemoveAll(dataDir)
}

func TestSavePost(t *testing.T) {
	PrepareTests()

	LoadPosts()
	SetPostManagerSettings(".json")
	NewPost(models.BlogPost{Content:"Post Content.", Author:"Auth"})
	post, _ := GetMostRecentPost()
	SavePost(post)
	postLoaded := LoadPostByPath("1.json")
	assert.True(t, postLoaded.PostID == "1")
	assert.True(t, postLoaded.Author == "Auth")
	assert.True(t, postLoaded.Content == "Post Content.")

	os.RemoveAll(dataDir)
}