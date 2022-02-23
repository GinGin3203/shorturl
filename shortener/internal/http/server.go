package http

import (
	"context"
	"fmt"
	"github.com/GinGin3203/shorturl/pkg/applogger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type shortenerService interface {
	GetOriginalURL(ctx context.Context, encoding string) (string, error)
	ShortenURL(ctx context.Context, encoding string) (string, error)
}

type server struct {
	Router chi.Router
	log    applogger.Logger
	srv    shortenerService
}

func NewServer(log applogger.Logger, service shortenerService) *server {
	s := &server{}
	s.initRoutes()
	return &server{
		log: log,
		srv: service,
	}
}

func (s *server) Run(port int) error {
	s.log.Infof("starting server on port: %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.Router)
}
