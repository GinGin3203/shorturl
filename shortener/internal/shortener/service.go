package shortener

import (
	"context"
	"github.com/GinGin3203/shorturl/pkg/applogger"
	"github.com/GinGin3203/shorturl/pkg/base63"
	"github.com/GinGin3203/shorturl/shortener/config"
	"github.com/GinGin3203/shorturl/shortener/internal/service_errors"
	"net/url"
	"strings"
)

type URLRepository interface {
	GetURLByID(ctx context.Context, id int) (string, error)
	InsertAndGetID(ctx context.Context, url string) (int, error)
}

type service struct {
	log  applogger.Logger
	repo URLRepository
	cfg  *config.Service
}

func NewService(log applogger.Logger, r URLRepository, cfg *config.Service) *service {
	return &service{
		log:  log,
		repo: r,
		cfg:  cfg,
	}
}

func (s *service) GetOriginalURL(ctx context.Context, shortUrl string) (string, error) {
	parsed, err := url.ParseRequestURI(shortUrl)
	if err != nil {
		s.log.Debug(err, shortUrl)
		return "", service_errors.ErrInvalidURL
	}
	enc := strings.TrimPrefix(parsed.Path, "/")
	if len(enc) < s.cfg.MinURLEncodingLength {
		s.log.Debug(shortUrl)
		return "", service_errors.ErrInvalidEncoding
	}
	id, err := base63.Decode(enc)
	if err != nil {
		s.log.Debug(err)
		return "", err
	}
	s.log.Debugf("new request: %d", id)
	origUrl, err := s.repo.GetURLByID(ctx, id)
	if err != nil {
		s.log.Debug(err)
		return "", err
	}

	return origUrl, nil
}

func (s *service) CreateAndGetShortURL(ctx context.Context, longUrl string) (string, error) {
	if _, err := url.ParseRequestURI(longUrl); err != nil {
		s.log.Debug(err)
		return "", service_errors.ErrInvalidURL
	}
	s.log.Debug("new request: %s", longUrl)
	id, err := s.repo.InsertAndGetID(ctx, longUrl)
	if err != nil {
		s.log.Error(err, id)
		return "", err
	}
	shortUrl := "https://" +
		s.cfg.ShortURLDomainName + "/" +
		base63.Encode(id, s.cfg.MinURLEncodingLength)
	return shortUrl, nil
}
