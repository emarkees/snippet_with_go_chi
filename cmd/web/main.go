package main

import (
	"flag"
	"github.com/emarkees/chi/internal/app"
	"github.com/emarkees/chi/internal/routes"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	/*
		Initialize a new instance of our application struct, containing the
		dependencies.
	*/

	app := &app.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	// r := chi.NewRouter()

	router := routes.SetUpRoutes(app)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  router,
	}

	infoLog.Printf("Server is running on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
