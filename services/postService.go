/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"de/vorlesung/projekt/2416160-5836402/models"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
	"errors"
	"log"
	"path/filepath"
)

var postSuffix string
var posts []models.BlogPost
var postsLoaded = false
var lastPostID = 0

func SetPostManagerSettings(postSuf string) {
	postSuffix = postSuf
}

func GetMostRecentPost() (post models.BlogPost, err error) {
	if !postsLoaded {
		log.Println("No Posts loaded.")
		return models.BlogPost{}, errors.New("no posts")
	}

	if len(posts) == 0 {
		return models.BlogPost{PostID:"0", Content:"No Posts."}, nil
	}

	mrPost := posts[0]
	for i := 1; i < len(posts); i++ {
		if posts[i].Date.After(mrPost.Date) {
			mrPost = posts[i]
		}
	}
	return mrPost, nil
}

func GetAllPosts() (retPosts []models.BlogPost, err error) {
	if !postsLoaded {
		log.Println("No Posts loaded.")
		return []models.BlogPost{}, errors.New("no posts loaded")
	}

	if len(posts) == 0 {
		return []models.BlogPost{{PostID:"0", Content:"No Posts."}}, nil
	}

	//TODO Sort posts by time (reverse list)

	return posts, nil
}

func GetAllPostsFromUser(username string) (retPosts []models.BlogPost, err error) {
	if !postsLoaded {
		log.Println("No Posts loaded.")
		return []models.BlogPost{}, errors.New("no posts loaded")
	}

	if len(posts) == 0 {
		return []models.BlogPost{{PostID:"0", Content:"No Posts."}}, nil
	}

	var filteredPosts []models.BlogPost
	for i, p := range posts {
		if p.Author == username {
			filteredPosts = append(filteredPosts, posts[i])
		}
	}

	if len(filteredPosts) == 0 {
		return []models.BlogPost{{PostID:"0", Content:"No Posts."}}, nil
	}

	return filteredPosts, nil
}

func LoadPosts() {
	if postsLoaded {
		log.Println("Posts already loaded.")
		return
	}

	files, err := ioutil.ReadDir(filepath.Join(dataDir, PostsDir))
	if err != nil {
		log.Println("No Posts found.")
	}

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), postSuffix) {
			posts = append(posts, LoadPostByPath(f.Name()))
		}
	}

	for _, p := range posts {
		idNum, err := strconv.Atoi(p.PostID)
		if err != nil {
			log.Println("Invalid PostID.")

		} else {
			if lastPostID < idNum {
				lastPostID = idNum
			}
		}
	}

	postsLoaded = true
}

func LoadPostByPath(filename string) models.BlogPost {
	post := models.BlogPost{}

	err := ReadJsonFile(PostsDir, filename[:len(filename)-len(postSuffix)], &post)
	if err != nil {
		log.Println(err)
		return post
	}

	return post
}

func SavePost(post models.BlogPost) {
	err := WriteJsonFile(PostsDir, post.PostID, post)
	if err != nil {
		log.Println("Could not write to Post-File.")
	}
}

func NewPost(post models.BlogPost) {
	lastPostID++
	post.PostID = strconv.Itoa(lastPostID)
	post.Date = time.Now()
	posts = append(posts, post)

	SavePost(post)
}

func AppendCommentToPost(postID string, comment *models.Comment) {
	for i, post := range posts {
		if post.PostID == postID {
			posts[i].Comments = append(post.Comments, *comment)

			SavePost(post)
			return
		}
	}
	log.Println("Post", postID, "not found.")
}

func DeletePost(postID string) {
	for i, post := range posts {
		if post.PostID == postID {
			posts = append(posts[:i], posts[i+1:]...)
			deleteDataFile(PostsDir, postID + postSuffix)
			return
		}
	}
	log.Println("Post", postID, "not found.")
}

func ChangePost(postID, content string) {
	for i, post := range posts {
		if post.PostID == postID {
			posts[i].Content = content
			SavePost(posts[i])
			return
		}
	}
	log.Println("Post", postID, "not found.")
}