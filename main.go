package main

import (
	"net/http"
	"os"

	"github.com/apoclyps/atlas/api/handlers"
	"github.com/apoclyps/atlas/api/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	path := getenv("GEOIP_CITY_DBPATH", "GeoLite2-City.mmdb")

	db := storage.NewIPDatabase(path)
	defer db.Lookup.Close()

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	lookup := handlers.NewIPLookupHandler(db)
	r.Get("/", lookup.Retrieve)
	r.Get("/ip/{ip}", lookup.Retrieve)

	http.ListenAndServe(":8080", r)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
