package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/emarkees/chi/internal/routes"
	"os"
)

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFOR\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	r := chi.NewRouter()

	// File is serve through the http.FileServer
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	routes.SetUpRoutes(r)

	infoLog.Printf("Server is running on %s", *addr)
	err := http.ListenAndServe(*addr, r)
	errorLog.Fatal(err)
}