package routing

import (
	"github.com/go-chi/chi/v5"
	"learn/go/internal/api/handlers"
	"net/http"
)

func Route() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Привет, мир!"))
		if err != nil {
			return
		}
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", handlers.Register)
		})
	})

	return r
}
