### Atlas

> A simple GeoIP lookup service written in Golang.


### Download the free database

```sh
wget http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz -O GeoLite2-City.tar.gz && tar -zxvf GeoLite2-City.tar.gz -C tmp/ --strip 1 && rm GeoLite2-City.tar.gz
```

### Build and run the service

```sh
docker-compose build service
docker-compose run --service-ports service
```

#### Querying the service

```sh
curl -s "http://http://localhost:8080/ip/197.135.113.222" | jq .
```