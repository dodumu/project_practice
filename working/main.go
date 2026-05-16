package main

import (
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

type User struct {
	Name     string
	Email    string
	Password string
	Courses  []string
}

func BaseHandler(w http.ResponseWriter, page string, data User) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/"+page,
		"templates/parts/footer.html",
		"templates/parts/navbar.html",
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		pass := r.FormValue("password")

		user = User{
			Name:     "David",
			Email:    email,
			Password: pass,
		}
		if user.Email != "davidrobertq2@gmail.com" && user.Password != "49321111@David" {
			http.Error(w, "incorrect email or password", http.StatusBadRequest)
			return
		}
		values := url.Values{}
		values.Add("email", email)
		values.Add("Password", pass)
		values.Add("name", user.Name)
		if user.Email == "davidrobertq2@gmail.com" && user.Password == "49321111@David" {
			redirectTo := "/dashboard?" + values.Encode()
			http.Redirect(w, r, redirectTo, http.StatusSeeOther)
			return
		}

	}
	BaseHandler(w, "login.html", user)
}

func DashHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	name := r.URL.Query().Get("name")

	if r.Method == http.MethodPost {
		Go := r.FormValue("Go")
		HTML := r.FormValue("HTML")
		Base := r.FormValue("Databases")
		Algo := r.FormValue("Algorithms")
		Css := r.FormValue("Css")
		Dev := r.FormValue("DevOps")

		if Go == "on" {
			Go = "golang"
		}
		if HTML == "on" {
			HTML = "HTML"
		}
		if Base == "on" {
			Base = "Database"
		}
		if Algo == "on" {
			Algo = "Algorithim"
		}
		if Css == "on" {
			Css = "Css"
		}
		if Dev == "on" {
			Dev = "DevOps"
		}
		courses := []string{
			Go,
			HTML,
			Base,
			Algo,
			Css,
			Dev,
		}

		var selectedCourses []string
		for _, course := range courses {
			if course != "" {
				selectedCourses = append(selectedCourses, course)
			}
		}

		joinedCourses := strings.Join(selectedCourses, ",")

		params := url.Values{}
		params.Add("email", email)
		params.Add("name", name)
		params.Add("Courses", joinedCourses)

		http.Redirect(w, r, "/profile?"+params.Encode(), http.StatusSeeOther)
		return
	}

	data := User{Name: name, Email: email}
	BaseHandler(w, "dash.html", data)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	courses := r.URL.Query().Get("Courses")

	course := strings.Split(courses, ",")
	var data User
	for _, val := range course {
		data.Courses = append(data.Courses, val)

	}
	BaseHandler(w, "profile.html", data)
}

func main() {
	http.HandleFunc("/", LoginHandler)
	http.HandleFunc("/dashboard", DashHandler)
	http.HandleFunc("/profile", ProfileHandler)

	http.ListenAndServe(":8099", nil)
}
