package store

import (
	"fmt"
	"net/http"
	"sync"
)

type smartUrlSet map[*SmartUrl]bool

type Memory struct {
	urls map[string]smartUrlSet
}

func NewMemory() *Memory {
	return &Memory{
		urls: make(map[string]smartUrlSet),
	}
}

func (m *Memory) GetUrl(key string, hour int) (*SmartUrl, error) {
	urls, ok := m.urls[key]
	if !ok {
		return nil, fmt.Errorf("url not found")
	}

	for url := range urls {
		if url.StartHour <= hour && url.EndHour > hour {
			return url, nil
		}
	}

	return nil, fmt.Errorf("url not found")
}

func (m *Memory) SetUrls(key string, urls []*SmartUrl) error {
	if _, ok := m.urls[key]; ok {
		return fmt.Errorf("key already exists")
	}

	m.urls[key] = make(smartUrlSet, len(urls))
	for _, url := range urls {
		m.urls[key][url] = true
	}

	return nil
}

func (m *Memory) RefreshUrls(key string) error {
	urls, ok := m.urls[key]
	if !ok {
		return fmt.Errorf("url not found")
	}

	var wg sync.WaitGroup
	urlsToBeRemoved := make(chan *SmartUrl, len(urls))
	for url := range urls {
		wg.Add(1)
		go func(url *SmartUrl, c chan<- *SmartUrl) {
			defer wg.Done()
			resp, err := http.Get(url.Url)
			if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 400 {
				c <- url
			}
		}(url, urlsToBeRemoved)
	}
	wg.Wait()

	close(urlsToBeRemoved)
	for url := range urlsToBeRemoved {
		delete(urls, url)
	}

	return nil
}
