/*
    Matrikelnummern: 5836402, 2416160
*/

package main

import (
	"net/http"
	"regexp"
	"de/vorlesung/projekt/2416160-5836402/handlers"
	"de/vorlesung/projekt/2416160-5836402/services"
	"de/vorlesung/projekt/2416160-5836402/global"
)

var validPath = regexp.MustCompile("^/(edit|save|view|login)/([a-zA-Z0-9]*)$")

func main() {
	e := services.LoadSettings()
	if e != nil {
		return
	}
	e = services.LoadUsers()
	if e != nil {
		return
	}

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/view/", makeResourceHandler(handlers.ViewHandler))
	http.HandleFunc("/login/", makeHandler(handlers.LoginHandler))
	http.ListenAndServeTLS(":" + global.Settings.PortNumber, global.Settings.CertFile, global.Settings.KeyFile, nil)
}

func makeResourceHandler(fn func(w http.ResponseWriter, r *http.Request, id string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.Redirect(w, r, "/index", http.StatusFound)
			return
		}
		fn(w, r, m[2])
	}
}

func makeHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.Redirect(w, r, "/index", http.StatusFound)
			return
		}
		if m[2] != "" {
			http.Redirect(w, r, m[1], http.StatusFound)
			return
		}
		fn(w, r)
	}
}