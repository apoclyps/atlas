package models

type IPConfig struct {
	Addr string
}

type IP struct {
	IP         string  `json:"ip"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	CountryISO string  `json:"country-iso"`
	Continent  string  `json:"continent"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"lng"`
	TimeZone   string  `json:"time-zone"`
}
