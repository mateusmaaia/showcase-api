# Showcase-api

Showcase-api is an Go RESTful Api for dealing with sale and rent of real estates of multiple ventures.

## Dependencies and Technologies used

- __[Go 1.11](https://golang.org/doc/go1.11)__ - This project uses [go modules](https://tip.golang.org/doc/go1.11#modules) as a package manager, so you need to use at least [Go 1.11](https://golang.org/doc/go1.11) for development mode.
- __[Docker](https://docs.docker.com)__ - Docker is a platform built for developers to build and run applications.
- __[Gin](https://github.com/gin-gonic/gin)__ - Gin is a HTTP web framework written in Go (Golang).
- __[Testify](https://github.com/gin-gonic/gin)__ - A toolkit with common assertions and mocks that plays nicely with the standard library.
- __[Godotenv](https://github.com/joho/godotenv)__ - A Go port of Ruby's dotenv library (Loads environment variables from `.env`.).

## Installation

**Development mode**
```bash
mv .example.env .env
go run main.go
```

**Production mode** - _I highly encourage the usage of any CICD feature (Eg. CircleCI) and any image repository (Eg. ECR)_

Run once to create the image locally:
```bash
docker build -t [image_name]:[version/tag] .
```

Run to start a container with your image:
```bash
docker run -p 8090:8090 --name=[container_name] -d  [image_name]
```

## Tests

To run tests just run:
```bash
go test ./... -cover
```

_I highly encourage running the tests on CICD pipeline_

## Usage

We have one main endpoint to retrieve real estates information and a health check:

### Health Check
`GET /health`
```JSON
"All engines running. Liftoff!"
```

### Real Estates
`GET /real-estates/[venture]?pageSize=[N]&pageNumber=[N]`

As you can see, we have 3 variables:

- __Venture__: That can be: **zap** or **viva-real**
- __PageNumber__: That is an integer
- __PageSize__: That is an integer

Each request have a metadata and our response for real estates will be inside listings, as you can se below:

```JSON
"pageNumber":1,
"pageSize":2,
"totalCount":977,
"listings": [...]
```
### To Do
- [X] Coding the challenge
- [X] Documentation
- [ ] Writing tests
- [X] Fixing docker image




