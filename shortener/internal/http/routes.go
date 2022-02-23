package http

import "github.com/go-chi/chi/v5"

func (s *server) initRoutes() {
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/shorturl", s.getOriginalURL)
		r.Post("/shorturl", s.generateShortURL)
	})

	s.Router = r
}
