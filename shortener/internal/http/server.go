package http

import (
	"context"
	"fmt"
	"github.com/GinGin3203/shorturl/pkg/applogger"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

type shortenerService interface {
	GetOriginalURL(ctx context.Context, shortUrl string) (string, error)
	CreateAndGetShortURL(ctx context.Context, longUrl string) (string, error)
}

type server struct {
	router  chi.Router
	log     applogger.Logger
	service shortenerService
}

func NewServer(log applogger.Logger, service shortenerService) *server {
	s := &server{
		log:     log,
		service: service,
	}
	s.initRoutes()
	return s
}

func (s *server) Run(port int, requestTimeoutSeconds int) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.TimeoutHandler(s.router, time.Duration(requestTimeoutSeconds)*time.Second, "Request Timeout"),
		// по-хорошему эти значения тоже надо выносить в конфиг
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       20 * time.Second,
	}

	return srv.ListenAndServe()
}
