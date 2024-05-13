package main

import (
	"log"
	"net/http"

	"troncy.fr/go-webserver-template/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/hello", handlers.HandleHello)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Démarrage du serveur sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
