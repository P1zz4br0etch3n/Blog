/*
    Matrikelnummern: 5836402, 2416160
*/

package handlers

import (
	"net/http"
	"de/vorlesung/projekt/2416160-5836402/services"
	"path/filepath"
	"html/template"
	"de/vorlesung/projekt/2416160-5836402/pages"
	"log"
	"de/vorlesung/projekt/2416160-5836402/models"
)

var templates = template.Must(template.ParseFiles(
	filepath.Join("tmpl", "index.html"),
	filepath.Join("tmpl", "login.html"),
	filepath.Join("tmpl", "view.html"),
))

func renderTemplate(w http.ResponseWriter, tmpl string, page interface{}) {
	e := templates.ExecuteTemplate(w, tmpl+".html", page)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := pages.IndexPage{
		Title: "Welcome to our Blog!",
	}

	session, err := services.CheckSession(r)
	if err != nil {
		renderTemplate(w, "index", page)
		return
	}

	page.UserLoggedIn = true
	page.UserName = session.UserName

	renderTemplate(w, "index", page)
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

func ViewHandler(w http.ResponseWriter, r *http.Request, id string) {
	var b models.Blog
	e := services.ReadJsonFile(id, services.BlogDir, &b)
	if e != nil {
		log.Println(e.Error())
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	renderTemplate(w, "view", b)
}
