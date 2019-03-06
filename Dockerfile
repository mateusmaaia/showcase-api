FROM golang:1.12-stretch AS build-env

LABEL maintainer "mateusinocenciomaia@gmail.com"

RUN mkdir /showcase-api
WORKDIR /showcase-api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN mv ./.example.env ./env

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/showcase-api

EXPOSE 8090
ENTRYPOINT ["/go/bin/showcase-api"]