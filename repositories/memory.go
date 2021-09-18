package repositories

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/RotemWald/smart-short-link/entities"
)

type Memory struct {
	urls map[string]smartUrlSet
	mu   sync.RWMutex
}

func NewMemory() *Memory {
	return &Memory{
		urls: make(map[string]smartUrlSet),
	}
}

func (m *Memory) GetUrl(key string, hour int) (*entities.SmartUrl, error) {
	m.mu.RLock()
	urls, ok := m.urls[key]
	m.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("url not found")
	}

	for url := range urls {
		if url.StartHour <= hour && url.EndHour > hour {
			return url, nil
		}
	}
	for url := range urls {
		return url, nil // default url in case there was not matched url found in the above loop
	}

	return nil, fmt.Errorf("url not found")
}

func (m *Memory) SetUrls(key string, urls []*entities.SmartUrl) error {
	m.mu.RLock()
	if _, ok := m.urls[key]; ok {
		return fmt.Errorf("key already exists")
	}
	m.mu.RUnlock()

	m.mu.Lock()
	m.urls[key] = make(smartUrlSet, len(urls))
	for _, url := range urls {
		m.urls[key][url] = true
	}
	m.mu.Unlock()

	return nil
}

func (m *Memory) RefreshUrls(key string) error {
	m.mu.RLock()
	urls, ok := m.urls[key]
	if !ok {
		return fmt.Errorf("url not found")
	}
	m.mu.RUnlock()

	var wg sync.WaitGroup
	brokenUrls := make(chan *entities.SmartUrl, len(urls))
	sem := make(chan struct{}, 8) // this is just an example for controlling the count of concurrent goroutines
	for url := range urls {
		wg.Add(1)
		sem <- struct{}{}
		go func(url *entities.SmartUrl, c chan<- *entities.SmartUrl) {
			defer func() {
				<-sem
				wg.Done()
			}()
			resp, err := http.Get(url.Url)
			if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 400 {
				c <- url
			}
		}(url, brokenUrls)
	}
	wg.Wait()

	close(brokenUrls) // channel can be safely closed as no one writes to the channel anymore at this time
	m.mu.Lock()
	for url := range brokenUrls {
		delete(urls, url)
	}
	m.mu.Unlock()

	return nil
}
