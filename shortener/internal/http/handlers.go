package http

import (
	"encoding/json"
	"errors"
	"github.com/GinGin3203/shorturl/shortener/internal/service_errors"
	"net/http"
)

// Response по-хорошему должен быть разбит на несколько структур
type Response struct {
	OriginalURL string `json:"original_url,omitempty"`
	ShortURL    string `json:"short_url,omitempty"`
	Error       string `json:"error,omitempty"`
}

func (s *server) getOriginalURL(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		s.log.Debug("url query expected: %v", r.URL.Query())
		s.writeResponse(w, http.StatusBadRequest, &Response{Error: "url query expected"})
		return
	}
	s.log.Debugf("request url: %s", url)
	origUrl, err := s.service.GetOriginalURL(r.Context(), url)
	if err != nil {
		switch {
		case errors.Is(err, service_errors.ErrInvalidURL):
			s.log.Debug(err, " ", url)
			s.writeResponse(w, http.StatusBadRequest, &Response{Error: err.Error()})
		case errors.Is(err, service_errors.ErrShortURLNotFound):
			s.log.Debug(err, " ", url)
			s.writeResponse(w, http.StatusNotFound, &Response{Error: err.Error()})
		default:
			s.log.Error(err, " ", url)
			s.writeResponse(w, http.StatusInternalServerError, &Response{Error: err.Error()})
		}
		return
	}
	s.log.Info("success ", url)
	s.writeResponse(w, http.StatusOK, &Response{OriginalURL: origUrl})
}

type Request struct {
	URL string `json:"url"`
}

func (s *server) createAndGetShortURL(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.log.Debugf("error when unmarshaling: %v", err)
		s.writeResponse(w, http.StatusBadRequest, &Response{Error: err.Error()})
		return
	}
	url := req.URL
	s.log.Debugf("request url: %s", url)
	shortUrl, err := s.service.CreateAndGetShortURL(r.Context(), url)
	if err != nil {
		switch {
		case errors.Is(err, service_errors.ErrInvalidURL):
			s.log.Debug(err)
			s.writeResponse(w, http.StatusBadRequest, &Response{Error: err.Error()})
		default:
			s.log.Error(err)
			s.writeResponse(w, http.StatusInternalServerError, &Response{Error: err.Error()})
		}
		return
	}
	s.log.Info("success ", url)
	s.writeResponse(w, http.StatusOK, &Response{ShortURL: shortUrl})
}

func (s *server) writeResponse(w http.ResponseWriter, code int, payload *Response) {
	r, err := json.Marshal(payload)
	if err != nil {
		s.log.Error("unable to unmarshal: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}
	w.WriteHeader(code)
	w.Write(r)
}
