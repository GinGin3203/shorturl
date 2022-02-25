package main

import (
	"context"
	"fmt"
	"github.com/GinGin3203/shorturl/pkg/applogger"
	"github.com/GinGin3203/shorturl/shortener/config"
	"github.com/GinGin3203/shorturl/shortener/internal/http"
	"github.com/GinGin3203/shorturl/shortener/internal/shortener"
	"github.com/GinGin3203/shorturl/shortener/internal/storage/memory"
	"github.com/GinGin3203/shorturl/shortener/internal/storage/postgres"
	"time"
)

func main() {
	config.Init()
	cfg := config.Get()
	log := applogger.New(cfg.Service.Name)
	defer log.Sync()

	var repository shortener.URLRepository
	var err error
	switch config.Get().Storage {
	case "memory":
		repository = memory.NewStorage()
	case "postgres":
		time.Sleep(10 * time.Second) // чтобы постгрес успел запуститься
		pgURL := config.Get().Postgres.URL
		repository, err = postgres.NewConn(context.Background(), pgURL)
		if err != nil {
			panic(fmt.Sprintf("(main) couldn't connect to postgres: %v", err))
		}
	default:
		panic(fmt.Sprintf("(main) unknown value: %s", config.Get().Storage))
	}

	service := shortener.NewService(
		log,
		repository,
		&config.Get().Service,
	)
	server := http.NewServer(
		log,
		service,
	)
	log.Infof("Starting server at port %d...", cfg.Http.Port)
	log.Fatal(server.Run(cfg.Http.Port, cfg.Http.RequestTimeoutSeconds))
}
