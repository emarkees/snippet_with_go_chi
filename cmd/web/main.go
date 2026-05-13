package main

import (
	"flag"
	"github.com/emarkees/chi/internal/routes"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFOR\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	/*
		Initialize a new instance of our application struct, containing the
		dependencies.
	*/

	app := &application.Application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	r := chi.NewRouter()

	// File is serve through the http.FileServer
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	routes.SetUpRoutes(r, app)

	srv := &http.Server{
		Addr:      *addr,
		ErrorLog: errorLog,
		Handlers: r,
	}

	infoLog.Printf("Server is running on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
