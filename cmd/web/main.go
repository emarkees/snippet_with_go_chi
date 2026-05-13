package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/emarkees/chi/internal/routes"
	"os"
)

/*
	Define an application struct to hold the application-wide dependencies for the
	web application. For now we'll only include fields for the two custom loggers, but
	we'll add more to it as the build progresses.
*/

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

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

	srv := &http.Server{
		Add:  *addr,
		ErrorLog: errorLog,
		handlers: r
	}

	infoLog.Printf("Server is running on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}