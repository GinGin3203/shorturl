package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *server) initRoutes() {
	r := chi.NewRouter()
	r.Route("/api/v1/shortener", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Get("/", s.getOriginalURL)
		r.Post("/", s.createAndGetShortURL)
	})

	s.router = r
}
