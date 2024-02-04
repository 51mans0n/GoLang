package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type PageVariables struct {
	Title string
}

type Character struct {
	Name     string
	Age      int
	Bankai   string
	ImageURL string
}

var characters = []Character{
	{"Ichigo Kurosaki", 15, "Tensa Zangetsu", "/source/IchigoKurosaki.jpg"},
	{"Rukia Kuchiki", 150, "Hakka no Togame", "/source/RukiaKuchiki.jpg"},
	{"Genryusai Yamamoto", 1500, "Zanka no Tachi", "/source/Yamamoto.jpg"},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pageVars := PageVariables{
		Title: "Bleach Manga",
	}

	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, pageVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func CharactersHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/characters.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, characters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/characters", CharactersHandler)
	http.ListenAndServe(":8080", r)
}
