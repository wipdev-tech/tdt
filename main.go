package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)

    case "POST":
        err := r.ParseForm()
        if err != nil {
            fmt.Println("feeh 7aga msh tamam")
        }

        fmt.Println(r.PostForm.Get("email"))
        fmt.Println(r.PostForm.Get("password"))

        http.Redirect(w, r, "/", http.StatusFound)
        
    default:
        fmt.Println("Unsupported method")
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
