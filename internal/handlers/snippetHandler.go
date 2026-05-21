package handlers

import (
	"fmt"
	// "html/template"
	"net/http"
	"strconv"
	"errors"

	"github.com/emarkees/chi/internal/app"
	"github.com/emarkees/chi/internal/models"
)

func Home(app *app.Application) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			notFound(w)
			return
		}

		snippets, err := app.Snippets.Latest()
		if err != nil {
			serverError(app, w, err)
			return
		}

		for _, snippet := range snippets {
			fmt.Fprintf(w, "%+v\n", snippet)
		}

		// files := []string{
		// 	"./ui/html/base.tmpl",
		// 	"./ui/html/partials/nav.tmpl",
		// 	"./ui/html/pages/home.tmpl",
		// }

		// ts, err := template.ParseFiles(files...)
		// if err != nil {
		// 	serverError(app, w, err)
		// 	return
		// }

		// err = ts.ExecuteTemplate(w, "base", nil)
		// if err != nil {
		// 	serverError(app, w, err)
		// 	return
		// }
	}
}

func CreateSnippet(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			clientError(w, http.StatusMethodNotAllowed)
			return
		}

		title := "O snail"
		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		expiresAt := 7

		id, err := app.Snippets.Insert(title, content, expiresAt)
		if err != nil {
			serverError(app, w, err)
			return
		}

		app.InfoLog.Printf("Snippet created with ID %d", id)
		http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)

		// w.WriteHeader(http.StatusCreated)

		// w.Write([]byte("Snippet created successfully"))
	}
}

func ViewSnippet(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			notFound(w)
			return
		}

		snippet, err := app.Snippets.Get(r.Context(), id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				notFound(w)
			} else {
				serverError(app, w, err)
			}
			return
		}

		fmt.Fprintf(w, "%+v", snippet)
	}
}



