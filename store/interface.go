package store

type SmartUrl struct {
	StartHour int    `json:"start_hour"`
	EndHour   int    `json:"end_hour"`
	Url       string `json:"url"`
}

type SmartUrlStore interface {
	GetUrl(key string, hour int) (*SmartUrl, error)
	SetUrls(key string, urls []*SmartUrl) error
	RefreshUrls(key string) error
}
