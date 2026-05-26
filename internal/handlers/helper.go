package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/emarkees/chi/internal/app"
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

func render(app *app.Application, w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.TemplateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		serverError(app, w, err)
		return
	}

	// Initialize a buffer
	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		serverError(app, w, err)
		return
	}

	w.WriteHeader(status)

	_, err = buf.WriteTo(w)
	if err != nil {
		serverError(app, w, err)
	}
}

func newTemplateData(app *app.Application, r *http.Request) *templateData {
    return &templateData{
        CurrentYear: time.Now().Year(),
    }
}
