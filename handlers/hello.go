package handlers

import (
	"html/template"
	"net/http"
)

type HelloPageData struct {
	Name string
	Age  string
}

func HandleHello(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	tmpl, err := template.ParseFiles("templates/hello.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := HelloPageData{
		Name: queryParams.Get("name"),
		Age:  queryParams.Get("age"),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
