package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)

    fmt.Println("Listening at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		db, err := sql.Open("sqlite3", "./tdt.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		rows, err := db.Query("SELECT * FROM users;")
		for rows.Next() {
			var id int
			var name string

			err = rows.Scan(&id, &name)
			if err != nil {
				panic(err)
			}

            fmt.Println(id, name)
		}

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
