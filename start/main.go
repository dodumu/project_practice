package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Email    string
	Password string
}

func BaseHandler(w http.ResponseWriter, page string, data User) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/"+page,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user = User{
			Email:    email,
			Password: password,
		}
		if user.Password == "David123" {
			http.Redirect(w, r, "/profile?email="+email, http.StatusSeeOther)
			return
		}
	}
	BaseHandler(w, "index.html", user)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	data := User{Email: email}
	BaseHandler(w, "login.html", data)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/profile", ProfileHandler)
	http.ListenAndServe(":8095", nil)
}
