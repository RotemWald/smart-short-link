package services

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/RotemWald/smart-short-link/entities"
	"github.com/RotemWald/smart-short-link/repositories"
	"github.com/google/uuid"
)

type SmartUrl struct {
	repository repositories.SmartUrl
	counter    uint64
	mu         map[string]time.Time
}

func NewSmartUrl(repository repositories.SmartUrl) *SmartUrl {
	return &SmartUrl{
		repository: repository,
		mu:         make(map[string]time.Time),
	}
}

func (s *SmartUrl) GetUrl(key string) (string, error) {
	url, err := s.repository.GetUrl(key, time.Now().UTC().Hour())
	if err != nil {
		return "", err
	}

	// this is an example for limiting the RefreshUrls method invocation per key on 1 instance application
	// here we limit invocation of specific key for a duration of 5 minutes
	// on a distributed environment we would have used Redis, for example, to maintain this mechanism
	then, ok := s.mu[key]
	if !ok {
		then = time.Now().UTC()
		s.mu[key] = then
	}
	now := time.Now().UTC()
	diff := now.Sub(then)
	if diff.Minutes() >= 5 {
		s.mu[key] = time.Now().UTC()
		go func(key string) {
			if err := s.repository.RefreshUrls(key); err != nil {
				fmt.Println(err)
			}
		}(key)
	}

	return url.Url, nil
}

func (s *SmartUrl) SetUrlsByUUID(urls []*entities.SmartUrl) (string, error) {
	u := uuid.New()
	return s.setUrls(u.String(), urls)
}

func (s *SmartUrl) SetUrlsByCounter(urls []*entities.SmartUrl) (string, error) {
	num := atomic.AddUint64(&s.counter, 1)
	key := fmt.Sprintf("a%s", strconv.FormatUint(num, 10))
	return s.setUrls(key, urls)
}

func (s *SmartUrl) setUrls(key string, urls []*entities.SmartUrl) (string, error) {
	if err := s.repository.SetUrls(key, urls); err != nil {
		return "", err
	}
	return key, nil
}
