FROM golang:1.12 AS build-env
EXPOSE 8090

LABEL maintainer "mateusinocenciomaia@gmail.com"

# Just to download dependencies and caching them
RUN mkdir /showcase-api
WORKDIR /showcase-api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/showcase-api

FROM alpine:latest
RUN apk add --update --no-cache ca-certificates git
WORKDIR /app
COPY --from=build-env /go/bin/showcase-api .
COPY --from=build-env /showcase-api/.example.env ./.env

CMD ["/app/showcase-api"]