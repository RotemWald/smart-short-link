package entities

type SmartUrl struct {
	StartHour int    `json:"start_hour"`
	EndHour   int    `json:"end_hour"`
	Url       string `json:"url"`
}
