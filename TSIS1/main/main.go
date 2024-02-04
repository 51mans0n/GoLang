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
	{"Ichigo Kurosaki", 15, "Tensa Zangetsu", "https://upload.wikimedia.org/wikipedia/en/1/1e/IchigoKurosakiBleach.jpg"},
	{"Rukia Kuchiki", 150, "Hakka no Togame", "https://upload.wikimedia.org/wikipedia/en/thumb/0/0c/RukiaKuchikiKubo.jpg/220px-RukiaKuchikiKubo.jpg"},
	{"Genryusai Yamamoto", 1500, "Zanka no Tachi", "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/b32f6482-8a64-4321-80e7-cf3df263f60b/ddp1ufa-a932e35e-b1ee-49f0-8655-5de2d7a01de7.jpg?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcL2IzMmY2NDgyLThhNjQtNDMyMS04MGU3LWNmM2RmMjYzZjYwYlwvZGRwMXVmYS1hOTMyZTM1ZS1iMWVlLTQ5ZjAtODY1NS01ZGUyZDdhMDFkZTcuanBnIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.hShH7xicsdutuykBSFG4jgTHTeEwLUmRlnujDcEL4go"},
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
