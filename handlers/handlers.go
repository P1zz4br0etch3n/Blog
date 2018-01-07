/*
    Matrikelnummern: 5836402, 2416160
*/

package handlers

import (
	"net/http"
	"de/vorlesung/projekt/2416160-5836402/services"
	"log"
	"de/vorlesung/projekt/2416160-5836402/models"
	"strings"
	"github.com/pkg/errors"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//Just for Fun
	query := r.URL.Query()
	msg := query.Get("sendDeveloperMessage")
	if msg != "" {
		log.Println(msg)
	}

	// get newest Post
	newestPost, err := services.GetMostRecentPost()

	// create model for template
	pageData := models.IndexPage{
		UserLoggedIn:    false,
		ShowArchiveLink: true,
		UserName:        "",
		Posts:           []models.BlogPost{newestPost},
	}

	//Check if user is logged in
	session, err := services.CheckSession(r)
	if err == nil {
		pageData.UserName = session.UserName
		pageData.UserLoggedIn = true
	}

	services.RenderTemplate(w, "index", pageData)
}

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	// read form data
	oldpwd := r.FormValue("oldpwd")
	newpwd1 := r.FormValue("newpwd1")
	newpwd2 := r.FormValue("newpwd2")

	//Get username from session
	var username string
	session, err := services.CheckSession(r)
	if err == nil {
		username = session.UserName
	} else {
		// redirect to index if no session
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// change password
	if newpwd1 == newpwd2 {
		err := services.ChangePassword(username, oldpwd, newpwd1)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		} else {
			err = services.ChangePassword(username, oldpwd, newpwd1)
		}
	} else {
		err = errors.New("the new passwords did not match")
	}

	if oldpwd == "" && newpwd1 == "" && newpwd2 == "" {
		err = nil
	}

	services.RenderTemplate(w, "chpass", err)
}

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	nickname := r.FormValue("nickname")
	comment := r.FormValue("comment")
	postID := r.FormValue("postID")

	//Prevent ugly empty names
	if strings.Trim(nickname, " ") == "" {
		nickname = "anonymous"
	} else {
		//User is not anonymous set nickname
		cookie := &http.Cookie{
			Name:   "nickname",
			Value:  nickname,
			MaxAge: 60 * 60 * 24 * 3650,
			Path:   "/",
			Secure: true,
		}
		http.SetCookie(w, cookie)
	}

	services.AppendCommentToPost(postID, &models.Comment{Content: comment, Nickname: nickname})

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	//Check if user is logged in
	_, err := services.CheckSession(r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	content := r.FormValue("content")
	author := r.FormValue("author")

	if content != "" && author != "" {
		services.NewPost(models.BlogPost{Author: author, Content: content})
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func ArchiveHandler(w http.ResponseWriter, r *http.Request) {
	allPosts, err := services.GetAllPosts()

	pageData := models.IndexPage{
		UserLoggedIn:    false,
		ShowArchiveLink: false,
		UserName:        "",
		Posts:           allPosts,
	}

	//Check if user is logged in
	session, err := services.CheckSession(r)
	if err == nil {
		pageData.UserLoggedIn = true
		pageData.UserName = session.UserName
	}

	services.RenderTemplate(w, "index", pageData)
}

func MyPostsHandler(w http.ResponseWriter, r *http.Request) {
	//Fetch name of the author
	var authorname string
	session, err := services.CheckSession(r)
	if err == nil {
		authorname = session.UserName
	} else {
		//Redirect if no session available
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	postsFromAuthor, err := services.GetAllPostsFromUser(authorname)

	pageData := models.IndexPage{
		UserLoggedIn:    true,
		ShowArchiveLink: true,
		UserName:        authorname,
		Posts:           postsFromAuthor,
	}

	services.RenderTemplate(w, "myposts", pageData)
}

func ChangePostHandler(w http.ResponseWriter, r *http.Request) {
	//Check if user is logged in
	_, err := services.CheckSession(r)
	if err != nil {
		//Redirect if no session available
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	postID := r.FormValue("postID")
	content := r.FormValue("content")
	action := r.FormValue("action")

	if action == "delete" {
		services.DeletePost(postID)

		http.Redirect(w, r, "/myposts", http.StatusTemporaryRedirect)
	} else if action == "edit" {
		services.ChangePost(postID, content)

		http.Redirect(w, r, "/myposts", http.StatusTemporaryRedirect)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == http.MethodGet {
		//If there is a session running redirect to index
		_, err := services.CheckSession(r)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		services.RenderTemplate(w, "login", nil)
	}

	if strings.ToUpper(r.Method) == http.MethodPost {
		//Validate user and redirect to index after successful login
		username := r.FormValue("username")
		password := r.FormValue("password")
		err := services.VerifyUser(username, password)
		if err == nil {
			cookie, err := services.GenerateSessionCookie()
			if err != nil {
				services.RenderTemplate(w, "login", err)
			}
			services.GenerateSession(username, cookie.Value)
			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		services.RenderTemplate(w, "login", err)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	services.DestroySession(r)

	//Redirect to index
	http.Redirect(w, r, "/", http.StatusFound)
}
