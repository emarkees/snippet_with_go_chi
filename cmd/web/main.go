package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/emarkees/chi/internal/app"
	"github.com/emarkees/chi/internal/routes"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "postgres://learn2earn:pass@localhost:5432/snippetbox", "Database connection string")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Connect to database (POOL)
	dbpool, err := pgxpool.New(context.Background(), *dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer dbpool.Close()

	// Test connection
	err = dbpool.Ping(context.Background())
	if err != nil {
		errorLog.Fatal(err)
	}

	// Application struct
	app := &app.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		// DB: dbpool, // (recommended to add this)
	}

	router := routes.SetUpRoutes(app)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  router,
	}

	infoLog.Printf("Server is running on %s", *addr)

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB()