package main

import (
	"net/http"
	"text/template"
)

type Student struct {
	Name string
}

func BaseHandler(w http.ResponseWriter, page string, data Student) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/"+page,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
}

func BoardHandler(w http.ResponseWriter, r *http.Request) {
	user := Student{
		Name: "David",
	}

	BaseHandler(w, "board.html", user)
}

func main() {
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.HandleFunc("/", BoardHandler)
	http.ListenAndServe(":8089", nil)
}
