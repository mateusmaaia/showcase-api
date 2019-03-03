package main

import (
	"os"

	_ "github.com/joho/godotenv"
	"github.com/mateusmaaia/showcase-api/api"
	"github.com/mateusmaaia/showcase-api/domains"
	"github.com/mateusmaaia/showcase-api/infrastructure"
	"github.com/mateusmaaia/showcase-api/repository"
	"github.com/mateusmaaia/showcase-api/services"
)

func main() {
	storeRealState := make(map[string][]domains.RealEstate)
	realEstateRepository := &repository.RealEstateRepository{StoreRealState: storeRealState}
	realEstateService := &services.RealEstateService{}

	server := &api.Server{}
	server.Run()
}

func seedRepositories() {
	sourceUrl := os.Getenv("S3_BUCKET")
	seed := &infrastructure.Seed{sourceURL: sourceURL}
	realEstates, err := seed.Import()

	if err != nil {
		panic(err)
	}

	for _, realEstate := range realEstates {
		services.RealEstateService.Insert(realEstate)
	}
}
