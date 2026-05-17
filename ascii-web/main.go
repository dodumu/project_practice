package main

import (
	"html/template"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type User struct {
	Result string
}

func BaseHandler(w http.ResponseWriter, page string, data User) {
	tmpl, err := template.ParseFiles(
		"templates/index.html",
		"templates/"+page,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := User{}

	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		font := r.FormValue("font")

		if text == "" {
			http.Error(w, "empty text", http.StatusBadRequest)
			return
		}
		result := Render(text, font)
		data.Result = result

		http.Redirect(w, r, "/ascii-art?result="+url.QueryEscape(result), http.StatusSeeOther)
	}
	BaseHandler(w, "home.html", data)
}

func ResultHander(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("result")

	user := User{
		Result: data,
	}

	BaseHandler(w, "result.html", user)
}

func Render(text string, banner string) string {
	data, err := os.ReadFile(banner)
	if err != nil {
		return "invalid file"
	}
	lines := strings.Split(string(data), "\n")
	font := make(map[rune][]string)

	for ch := ' '; ch <= '~'; ch++ {
		start := (int(ch) - 32) * 9
		font[ch] = lines[start+1 : start+9]
	}
	var output strings.Builder
	input := strings.Split(text, "\n")

	for _, word := range input {
		if word == "" {
			output.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range word {
				glyph, ok := font[ch]
				if ok {
					output.WriteString(glyph[row])

				}

			}
			output.WriteString("\n")
		}
	}
	return output.String()
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", ResultHander)

	http.ListenAndServe(":8056", nil)
}
