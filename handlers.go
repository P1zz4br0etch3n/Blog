/*
    Matrikelnummern: 5836402, 2416160
*/

package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	b, e1 := loadBlog(r.URL.Path[len("/view/"):])
	if e1 != nil {
		fmt.Print(e1)
	}

	t, e2 := template.ParseFiles("view.html")
	if e2 != nil {
		fmt.Print(e2)
	}

	t.Execute(w, b)
}