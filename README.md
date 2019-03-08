# Showcase-api
![Quality Gate Status](https://svgshare.com/i/Bde.svg)  
![Maintainability Rating](https://svgshare.com/i/Bdu.svg)
![Security Rating](https://svgshare.com/i/Beo.svg)
![Reliability](https://svgshare.com/i/Bez.svg)
![Lines of Code](https://svgshare.com/i/Bdv.svg)
![Bugs](https://svgshare.com/i/Beg.svg)
![Technical Debt](https://svgshare.com/i/Bdw.svg)
![Duplicated Lines](https://svgshare.com/i/BeL.svg)
![Vulnerability](https://svgshare.com/i/Bdf.svg)
![Code Smells](https://svgshare.com/i/Bep.svg)

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

## <a name="tests"></a>Tests

To run tests just run:
```bash
go test ./... -cover
```
PS: On those scripts I focused on testing all the business rules (100% coverage), not considering some side tests.

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

- __[Required] Venture__: That can be: **zap** or **viva-real**
- __[Optional] PageNumber__: That is an integer / **Standard value: 1**
- __[Optional] PageSize__: That is an integer / **Standard value: 50**

Each request have a metadata and our response for real estates will be inside listings, as you can se below:

```JSON
"pageNumber":1,
"pageSize":2,
"totalCount":977,
"listings": [...]
```

### SonarQube
I used SonarQube to get some statistics, unfortunely I don't have a servidor to use the live images, so I exported on a SVG Online share and added here. 

Anyway, I uploaded here the configurations that I Used, so you can check the statistics. 

Other problem that I got was to get test coverage, so I didn't uploaded that statistics :(, you can check using [Tests](#tests) command




