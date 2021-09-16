package storage

type SmartUrl struct {
	StartHour int    `json:"start_hour"`
	EndHour   int    `json:"end_hour"`
	Url       string `json:"url"`
}

type SmartUrlStorage interface {
	GetUrl(key string, hour int) (*SmartUrl, error)
	SetUrls(key string, urls []*SmartUrl) error
}
