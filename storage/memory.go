package storage

import "fmt"

type Memory struct {
	urls map[string][]*SmartUrl
}

func (m *Memory) GetUrl(key string, hour int) (*SmartUrl, error) {
	urls, ok := m.urls[key]
	if !ok {
		return nil, fmt.Errorf("url not found")
	}

	for _, url := range urls {
		if url.StartHour <= hour && url.EndHour >= hour {
			return url, nil
		}
	}

	return nil, fmt.Errorf("url not found")
}

func (m *Memory) SetUrls(key string, urls []*SmartUrl) error {
	if _, ok := m.urls[key]; ok {
		return fmt.Errorf("key already exists")
	}
	m.urls[key] = urls
	return nil
}
