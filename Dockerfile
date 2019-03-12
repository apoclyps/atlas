FROM golang:alpine as build

ENV GO111MODULE on
ENV GOLINT_VERSION 1.12.5

RUN mkdir -p $GOPATH/src/github.com/apoclyps/atlas/api && \
    mkdir -p /build && \
    apk --no-cache add curl git bash gcc libc-dev ca-certificates && \
    update-ca-certificates && \
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v$GOLINT_VERSION

COPY ./go.mod $GOPATH/src/github.com/apoclyps/atlas/api
COPY ./go.sum $GOPATH/src/github.com/apoclyps/atlas/api

WORKDIR $GOPATH/src/github.com/apoclyps/atlas/api

RUN go mod download

ADD . $GOPATH/src/github.com/apoclyps/atlas/api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /build/atlas-api .

FROM scratch
WORKDIR /root/
COPY db/GeoLite2-City.mmdb .
COPY --from=build /build/ .
CMD ["./atlas-api", "-port=8080"]
EXPOSE 8080
