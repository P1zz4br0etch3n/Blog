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
	"fmt"
	"bufio"
	"os"
	"strings"
	"de/vorlesung/projekt/WebBlog/settingsManager"
)

const version = "0.0.1 pre-alpha"
var validPath = regexp.MustCompile("^/(login|logout|chpass|comment|newpost|archive|myposts|change)/([a-zA-Z0-9]*)$")

func main() {
	e := services.LoadSettings()
	if e != nil {
		return
	}
	e = services.LoadUsers()
	if e != nil {
		return
	}
	services.SetPostManagerSettings(global.Settings.PostSuffix)
	services.LoadPosts()

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/login/", makeHandler(handlers.LoginHandler))
	http.HandleFunc("/logout/", makeHandler(handlers.LogoutHandler))
	http.HandleFunc("/chpass/", makeHandler(handlers.ChangePasswordHandler))
	//http.HandleFunc("/comment", handlers.CommentHandler)
	//http.HandleFunc("/newpost", handlers.NewPostHandler)
	//http.HandleFunc("/archive", handlers.ArchiveHandler)
	//http.HandleFunc("/myposts", handlers.MyPostsHandler)
	//http.HandleFunc("/change", handlers.ChangePostHandler)
	//go http.ListenAndServe(":80", http.HandlerFunc(handlers.TlsRedirect))
	go http.ListenAndServeTLS(":" + global.Settings.PortNumber, global.Settings.CertFile, global.Settings.KeyFile, nil)

	repl()
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
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		//if m[2] != "" {
		//	http.Redirect(w, r, m[1] + "/", http.StatusFound)
		//	return
		//}
		fn(w, r)
	}
}

func repl() {
	fmt.Println("\nStarting REPL")
	for true {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		cmd := strings.Trim(strings.Split(text, " ")[0], "\n")

		switch cmd {
		case "?":
			fmt.Println("q - quit\ns - settings\nv - version\nu - users\n")
		case "q":
			os.Exit(0)
		case "s":
			fmt.Println(global.Settings)
		case "v":
			fmt.Println(version)
		case "u":
			for _, value := range services.GetOnlineUserNames() {
				fmt.Println("User:", value)
			}
		case "r":
			if strings.Contains(text, "config") {
				settingsManager.ForceLoadSettingsFile()
			}
		default:
			if cmd != "" {
				fmt.Println(cmd, "is not a command")
			}
		}
	}
}