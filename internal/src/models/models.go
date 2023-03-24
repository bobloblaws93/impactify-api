package models

import (
	"time"
)

type PublisherInformation struct {
	Publisher   Publisher `json:"publisher"`
	Date        time.Time `json:"date"`
	Requests    int64     `json:"requests"`
	Revenue     float64   `json:"revenue"`
	Clicks      int64     `json:"clicks"`
	Impressions int64     `json:"impressions"`
}

type Publisher struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type PublisherTimeRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type Currency struct {
	ID     int64   `json:"id"`
	Symbol string  `json:"symbol"`
	Rate   float64 `json:"rate"`
}
