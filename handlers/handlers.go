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
	"errors"
	"de/vorlesung/projekt/2416160-5836402/global"
)

var templates = template.Must(template.ParseFiles(
	filepath.Join("tmpl", "index.html"),
	filepath.Join("tmpl", "login.html"),
	filepath.Join("tmpl", "view.html"),
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//If there is a session running redirect to index
		_, err := services.CheckSession(r)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		}
		renderTemplate(w, "login", nil)
	}

	if r.Method == http.MethodPost {
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


