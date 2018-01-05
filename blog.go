/*
    Matrikelnummern: 5836402, 2416160
*/

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}