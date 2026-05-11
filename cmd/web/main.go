package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/emarkees/chi/internal/routes"
)

func main() {
	r := chi.NewRouter()

	// File is serve through the http.FileServer
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	routes.SetUpRoutes(r)

	log.Println("Server is running on:8080")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}