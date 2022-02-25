package shortener

import (
	"context"
	"github.com/GinGin3203/shorturl/pkg/applogger"
	"github.com/GinGin3203/shorturl/pkg/base63"
	"github.com/GinGin3203/shorturl/shortener/config"
	"github.com/GinGin3203/shorturl/shortener/internal/service_errors"
	"github.com/GinGin3203/shorturl/shortener/internal/storage/memory"
	"testing"
)

func makeTestService() *service {
	return &service{
		log: applogger.New("test"), repo: memory.NewStorage(), cfg: &config.Service{
			Name:                 "url_service",
			MinURLEncodingLength: 5,
			ShortURLDomainName:   "o.co",
		},
	}
}

func TestServiceHappyPath(t *testing.T) {
	s := makeTestService()
	ctx := context.Background()
	origUrl := "https://yandex.ru?q=aboba"
	id, err := s.repo.InsertAndGetID(ctx, origUrl)
	if err != nil {
		t.Errorf("err not expected: %v", err)
	}
	shortUrl := "https://" + s.cfg.ShortURLDomainName + "/" + base63.Encode(id, s.cfg.MinURLEncodingLength)
	longUrl, err := s.GetOriginalURL(ctx, shortUrl)
	if err != nil {
		t.Errorf("err not expected: %v", err)
	}
	if longUrl != origUrl {
		t.Fatalf("expected %s to be equal %s", longUrl, origUrl)
	}
}

func TestService_GetOriginalURL(t *testing.T) {
	t.Run("short url not found if not put before", func(t *testing.T) {
		s := makeTestService()
		ctx := context.Background()
		_, err := s.GetOriginalURL(ctx, "https://o.co/AAAAA")
		if err != service_errors.ErrShortURLNotFound {
			t.Fatalf("expected err: %v, but got: %v", service_errors.ErrShortURLNotFound, err)
		}
	})
}

func TestService_InsertAndGetID(t *testing.T) {
	t.Run("same result if InsertAndGetID called twice", func(t *testing.T) {
		s := makeTestService()
		ctx := context.Background()
		origUrl := "https://yandex.ru?q=aboba"
		id1, err := s.repo.InsertAndGetID(ctx, origUrl)
		if err != nil {
			t.Errorf("err not expected: %v", err)
		}
		id2, err := s.repo.InsertAndGetID(ctx, origUrl)
		if err != nil {
			t.Errorf("err not expected: %v", err)
		}
		if id1 != id2 {
			t.Fatalf("expected %d to be equal %d", id1, id2)
		}
	})
}
