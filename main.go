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
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)

	case "POST":
		err := r.ParseForm()
		if err != nil {
			fmt.Println("feeh 7aga msh tamam")
		}

		formUsername := r.PostForm.Get("email")
		formPassword := r.PostForm.Get("password")

		db, err := sql.Open("sqlite3", "./tdt.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		var rowId int
		var rowUsername string
		var rowPassword string

		err = db.QueryRow(
			"SELECT * FROM users WHERE name = ?;",
			formUsername,
		).Scan(&rowId, &rowUsername, &rowPassword)

		if err != nil {
			panic(err)
		}

		if formPassword == rowPassword {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
            tmpl := template.Must(template.ParseFiles("templates/login.html"))
            tmpl.Execute(w, "Invalid credentials")
        }

	default:
		fmt.Println("Unsupported method")
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
