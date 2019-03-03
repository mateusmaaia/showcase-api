package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mateusmaaia/showcase-api/api"
	"github.com/mateusmaaia/showcase-api/controllers"
	"github.com/mateusmaaia/showcase-api/domains"
	"github.com/mateusmaaia/showcase-api/infrastructure"
	"github.com/mateusmaaia/showcase-api/repositories"
	"github.com/mateusmaaia/showcase-api/services"
)

func main() {
	storeRealEstates := make(map[string][]domains.RealEstate)
	realEstateRepository := repositories.RealEstateRepository{StoreRealEstates: storeRealEstates}
	realEstateService := services.RealEstateService{RealEstateRepository: realEstateRepository}
	realEstateController := controllers.RealEstateController{RealEstateService: realEstateService}

	healthController := controllers.HealthController{}

	seedRepositories(realEstateService)
	server := &api.Server{RealEstateController: realEstateController, HealthController: healthController}
	server.Run()
}

func seedRepositories(realEstateService services.RealEstateService) {
	sourceURL := os.Getenv("S3_BUCKET")
	seed := &infrastructure.Seed{SourceURL: sourceURL}
	realEstates, err := seed.Import()
	if err != nil {
		panic(err)
	}

	for _, realEstate := range realEstates {
		realEstateService.Insert(realEstate)
	}
}
