package http

import "net/http"

func (s *server) getOriginalURL(w http.ResponseWriter, r *http.Request) {
	s.log.Info("(getOriginalUrl) ")
}

func (s *server) generateShortURL(w http.ResponseWriter, r *http.Request) {
	s.log.Info("(getOriginalUrl) ")
}
