package main

import (
	"github.com/GinGin3203/shorturl/pkg/applogger"
	"github.com/GinGin3203/shorturl/shortener/config"
	"github.com/GinGin3203/shorturl/shortener/internal/http"
	shortener "github.com/GinGin3203/shorturl/shortener/internal/shortener"
)

func main() {
	cfg := config.Get()
	log := applogger.New(cfg.Service.Name)
	defer log.Sync()

	server := http.NewServer(
		log,
		shortener.NewService(
			log,
			nil,
		),
	)

	log.Fatal(server.Run(cfg.Http.Port))
}
