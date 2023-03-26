package models

type PublisherInformation struct {
	Publisher   Publisher `json:"publisher"`
	Impressions int64     `json:"impressions"`
	Requests    int64     `json:"requests"`
	Revenue     float64   `json:"revenue"`
	Clicks      int64     `json:"clicks"`
}

type Publisher struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type PublisherTimeRequest struct {
	StartDate string `json:"start_date" default:"2018-01-01"`
	EndDate   string `json:"end_date" default:"2020-01-01"`
}

type Currency struct {
	ID     int64   `json:"id"`
	Symbol string  `json:"symbol"`
	Rate   float64 `json:"rate"`
}
