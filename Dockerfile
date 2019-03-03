FROM golang:1.12 AS building-environment

LABEL maintainer "mateusinocenciomaia@gmail.com"

# Just to download dependencies and caching them
RUN mkdir /showcase-api
WORKDIR /showcase-api
COPY go.mod .
COPY go.sum .
COPY .example.env .env
RUN go mod download
COPY . .

# Building binary at a scratch image, so we can have the most lightweight
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/showcase-api
FROM scratch
COPY --from=building-environment /go/bin/showcase-api /go/bin/showcase-api
RUN apk add --update --no-cache ca-certificates git
ENTRYPOINT ["/go/bin/showcase-api"]