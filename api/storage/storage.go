package storage

import (
	"log"

	"github.com/oschwald/geoip2-golang"
)

// IPDatabase ...
type IPDatabase struct {
	Lookup *geoip2.Reader
}

// NewIPDatabase ...
func NewIPDatabase(path string) *IPDatabase {
	db, err := geoip2.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return &IPDatabase{Lookup: db}
}
