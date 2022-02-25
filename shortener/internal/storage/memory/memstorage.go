package memory

import (
	"context"
	"github.com/GinGin3203/shorturl/shortener/internal/service_errors"
	"sync"
)

type storage struct {
	storage        []string
	reverseStorage map[string]int
	sync.RWMutex
}

func NewStorage() *storage {
	s := &storage{}
	s.storage = make([]string, 0)
	s.reverseStorage = map[string]int{}
	return s
}

func (s *storage) GetURLByID(ctx context.Context, id int) (string, error) {
	s.RLock()
	defer s.RUnlock()
	if id >= len(s.storage) {
		return "", service_errors.ErrShortURLNotFound
	}

	return s.storage[id], nil
}

func (s *storage) InsertAndGetID(ctx context.Context, url string) (int, error) {
	s.RLock()
	u, ok := s.reverseStorage[url]
	s.RUnlock()
	if ok {
		return u, nil
	}
	s.Lock()
	defer s.Unlock()
	key := len(s.storage)
	s.storage = append(s.storage, url)
	s.reverseStorage[url] = key

	return key, nil
}
