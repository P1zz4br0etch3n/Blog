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
	e := templates.ExecuteTemplate(w, tmpl + ".html", page)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := pages.IndexPage{
		Title: "Welcome to our Blog!",
	}

	renderTemplate(w, "index", page)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "login", nil)
	}
	if r.Method == http.MethodPost {
		uname := r.FormValue("username")
		passwd := r.FormValue("password")
		e := services.AuthenticateUser(uname, passwd)
		if e != nil {
			log.Println(e.Error())
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		}
		http.Redirect(w, r, "/index", http.StatusFound)
	}
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