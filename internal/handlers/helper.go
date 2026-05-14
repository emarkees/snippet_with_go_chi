package handlers

import (
	"fmt"
	"github.com/emarkees/chi/internal/app"
	"net/http"
	"runtime/debug"
)

func serverError(app *app.Application, w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)

}

func notFound(w http.ResponseWriter) {
	clientError(w, http.StatusNotFound)
}
