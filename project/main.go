package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name string
}

func BaseHandler(w http.ResponseWriter, page string, data User) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		// "templates/parts/footer.html",
		// "templates/parts/navbar.html",
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

func DashHandler(w http.ResponseWriter, r *http.Request) {
	data := User{
		Name: "David",
	}
	BaseHandler(w, "dashboard.html", data)
}

func main() {
	fs := http.FileServer(http.Dir("style"))

	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.HandleFunc("/", DashHandler)
	http.ListenAndServe(":8082", nil)
}
