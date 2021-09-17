package store

import "github.com/RotemWald/smart-short-link/entities"

type SmartUrlStore interface {
	GetUrl(key string, hour int) (*entities.SmartUrl, error)
	SetUrls(key string, urls []*entities.SmartUrl) error
	RefreshUrls(key string) error
}
