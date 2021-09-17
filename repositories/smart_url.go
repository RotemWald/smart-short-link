package repositories

import "github.com/RotemWald/smart-short-link/entities"

type smartUrlSet map[*entities.SmartUrl]bool

type SmartUrl interface {
	GetUrl(key string, hour int) (*entities.SmartUrl, error)
	SetUrls(key string, urls []*entities.SmartUrl) error
	RefreshUrls(key string) error
}
