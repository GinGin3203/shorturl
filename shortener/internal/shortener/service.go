package service

import (
	"context"
	"github.com/GinGin3203/shorturl/pkg/applogger"
)

type ID int

type urlRepository interface {
	GetURLByID(ctx context.Context, id ID) (string, error)
	GetIDByURL(ctx context.Context, url string) (ID, error)
}

type service struct {
	log  applogger.Logger
	repo urlRepository
}

func NewService(log applogger.Logger, r urlRepository) *service {
	return &service{
		log:  log,
		repo: r,
	}
}

func (s *service) ShortenURL(ctx context.Context, longUrl string) (string, error) {
	panic("not implemented")
}

func (s *service) GetOriginalURL(ctx context.Context, shortUrl string) (string, error) {
	panic("not implemented")
}
