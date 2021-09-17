package services

import (
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/RotemWald/smart-short-link/entities"
	"github.com/RotemWald/smart-short-link/store"
	"github.com/google/uuid"
)

type SmartUrlService struct {
	store   store.SmartUrlStore
	counter uint64
}

func NewSmartUrlService(store store.SmartUrlStore) *SmartUrlService {
	return &SmartUrlService{
		store: store,
	}
}

func (s *SmartUrlService) SetUrlsByUuid(urls []*entities.SmartUrl) error {
	uuid := uuid.New()
	return s.setUrls(uuid.String(), urls)
}

func (s *SmartUrlService) SetUrlsByCounter(urls []*entities.SmartUrl) error {
	num := atomic.AddUint64(&s.counter, 1)
	str := fmt.Sprintf("a%s", strconv.FormatUint(num, 10))
	return s.setUrls(str, urls)
}

func (s *SmartUrlService) setUrls(key string, urls []*entities.SmartUrl) error {
	return nil
}
