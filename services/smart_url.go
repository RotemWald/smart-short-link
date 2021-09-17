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
}

func NewSmartUrl(repository repositories.SmartUrl) *SmartUrl {
	return &SmartUrl{
		repository: repository,
	}
}

func (s *SmartUrl) GetUrl(key string) (string, error) {
	url, err := s.repository.GetUrl(key, time.Now().UTC().Hour())
	if err != nil {
		return "", err
	}
	go func(key string) {
		s.repository.RefreshUrls(key)
	}(key)
	return url.Url, nil
}

func (s *SmartUrl) SetUrlsByUuid(urls []*entities.SmartUrl) (string, error) {
	key := uuid.New()
	return s.setUrls(key.String(), urls)
}

func (s *SmartUrl) SetUrlsByCounter(urls []*entities.SmartUrl) (string, error) {
	num := atomic.AddUint64(&s.counter, 1)
	str := fmt.Sprintf("a%s", strconv.FormatUint(num, 10))
	return s.setUrls(str, urls)
}

func (s *SmartUrl) setUrls(key string, urls []*entities.SmartUrl) (string, error) {
	if err := s.repository.SetUrls(key, urls); err != nil {
		return "", err
	}
	return key, nil
}
