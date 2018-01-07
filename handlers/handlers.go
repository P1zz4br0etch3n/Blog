/*
    Matrikelnummern: 5836402, 2416160
*/

package handlers

import (
	"net/http"
	"de/vorlesung/projekt/2416160-5836402/services"
	"path/filepath"
	"html/template"
	"log"
	"de/vorlesung/projekt/2416160-5836402/pages"
	"de/vorlesung/projekt/2416160-5836402/models"
	"strings"
	"de/vorlesung/projekt/2416160-5836402/global"
	"github.com/pkg/errors"
)

var templates = template.Must(template.ParseFiles(
	filepath.Join("tmpl", "index.html"),
	filepath.Join("tmpl", "login.html"),
	filepath.Join("tmpl", "myposts.html"),
	filepath.Join("tmpl", "chpass.html"),
))

func renderTemplate(w http.ResponseWriter, tmpl string, page interface{}) {
	e := templates.ExecuteTemplate(w, tmpl+".html", page)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//Just for Fun
	query := r.URL.Query()
	msg := query.Get("sendDeveloperMessage")
	if msg != "" {
		log.Println(msg)
	}

	newestPost, err := services.GetMostRecentPost()

	pageData := pages.IndexPage{
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

	renderTemplate(w, "index", pageData)
}

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	oldpwd := r.FormValue("oldpwd")
	newpwd1 := r.FormValue("newpwd1")
	newpwd2 := r.FormValue("newpwd2")

	//Get username from session
	var username string
	session, err := services.CheckSession(r)
	if err == nil {
		username = session.UserName
	} else {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	if newpwd1 == newpwd2 {
		err := services.ChangePassword(username, oldpwd, newpwd1)
		if err == nil {
			host := strings.Split(r.Host, ":")[0]
			target := "https://" + host + ":" + global.Settings.PortNumber
			http.Redirect(w, r, target, http.StatusTemporaryRedirect)
		} else {
			err = services.ChangePassword(username, oldpwd, newpwd1)
		}
	} else {
		err = errors.New("the new passwords did not match")
	}

	if oldpwd == "" && newpwd1 == "" && newpwd2 == "" {
		err = nil
	}

	renderTemplate(w, "chpass", err)
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

	//TODO direct to archive if user comes from archive

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
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

	pageData := pages.IndexPage{
		UserLoggedIn:    false,
		ShowArchiveLink: false,
		UserName:        "",
		Posts:           allPosts,
	}

	session, err := services.CheckSession(r)
	if err == nil {
		pageData.UserLoggedIn = true
		pageData.UserName = session.UserName
	}

	renderTemplate(w, "index", pageData)
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

	pageData := pages.IndexPage{
		UserLoggedIn:    true,
		ShowArchiveLink: true,
		UserName:        authorname,
		Posts:           postsFromAuthor,
	}

	renderTemplate(w, "myposts", pageData)
}

func ChangePostHandler(w http.ResponseWriter, r *http.Request) {
	_, err := services.CheckSession(r)
	if err != nil {
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
		}
		renderTemplate(w, "login", nil)
	}

	if strings.ToUpper(r.Method) == http.MethodPost {
		//Validate user and redirect to index after successful login
		username := r.FormValue("username")
		password := r.FormValue("password")
		err := services.VerifyUser(username, password)
		if err == nil {
			cookie, err := services.GenerateCookie()
			if err != nil {
				renderTemplate(w, "login", err)
			}
			services.GenerateSession(username, cookie.Value)
			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/", http.StatusFound)
		}
		renderTemplate(w, "login", err)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	services.DestroySession(r)

	//Redirect to index
	http.Redirect(w, r, "/", http.StatusFound)
}
