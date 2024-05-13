package handlers

import (
	"html/template"
	"net/http"
)

type HomePageData struct {
	Title string
	Body  string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := HomePageData{
		Title: "Page d'accueil",
		Body:  "Bienvenue sur notre site web !",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
