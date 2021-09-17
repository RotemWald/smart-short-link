package repositories

import (
	"github.com/RotemWald/smart-short-link/entities"
)

type Dummy struct{}

func NewDummy() *Dummy {
	return &Dummy{}
}

func (d *Dummy) GetUrl(key string, hour int) (*entities.SmartUrl, error) {
	return &entities.SmartUrl{
		StartHour: 0,
		EndHour:   23,
		Url:       "http://www.ynet.co.il",
	}, nil
}

func (d *Dummy) SetUrls(key string, urls []*entities.SmartUrl) error {
	return nil
}

func (d *Dummy) RefreshUrls(key string) error {
	return nil
}
