package handlers

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/apoclyps/atlas/api/models"
	"github.com/apoclyps/atlas/api/storage"
	"github.com/go-chi/chi"
)

// IPLookup ...
type IPLookup struct {
	db storage.IPDatabase
}

// NewIPLookupHandler ...
func NewIPLookupHandler(db *storage.IPDatabase) *IPLookup {
	return &IPLookup{db: *db}
}

// Retrieve a new ip lookup
func (p *IPLookup) Retrieve(w http.ResponseWriter, r *http.Request) {
	ip := chi.URLParam(r, "ip")
	parsedIP := net.ParseIP(ip)

	location, err := p.db.Lookup.City(parsedIP)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	data := &models.IP{
		IP:         ip,
		City:       location.City.Names["en"],
		Country:    location.Country.Names["en"],
		CountryISO: location.Country.IsoCode,
		Continent:  location.Continent.Code,
		Latitude:   location.Location.Latitude,
		Longitude:  location.Location.Longitude,
		TimeZone:   location.Location.TimeZone,
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, &data)
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
