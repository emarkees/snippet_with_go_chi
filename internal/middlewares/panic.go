package middlewares

import (
	"fmt"
	"net/http"

	"github.com/emarkees/chi/internal/app"
	"github.com/emarkees/chi/internal/handlers/helper"

)

func RecoverPanic(app *app.Application) func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.Header().Set("Connection",
					"close",
				)
				serverError(app, w,fmt.Errorf("%s", err))
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}