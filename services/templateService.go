/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"path/filepath"
	"html/template"
	"net/http"
	"log"
)

var templates *template.Template

func LoadTemplates() error {
	tmp, err := template.ParseFiles(
		filepath.Join("tmpl", "index.html"),
		filepath.Join("tmpl", "login.html"),
		filepath.Join("tmpl", "myposts.html"),
		filepath.Join("tmpl", "chpass.html"),
	)
	templates = tmp
	log.Println(err)
	return err
}

func RenderTemplate(w http.ResponseWriter, tmpl string, page interface{}) {
	e := templates.ExecuteTemplate(w, tmpl+".html", page)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
