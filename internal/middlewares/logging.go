package middlewares

import (
	"net/http"
	"github.com/emarkees/chi/internal/app"
)

func LogRequest(app *app.Application) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			app.InfoLog.Printf(
				"%s - %s %s %s",
				r.RemoteAddr,
				r.Proto,
				r.Method,
				r.URL.RequestURI(),
			)

			next.ServeHTTP(w, r)
		})
	}
}