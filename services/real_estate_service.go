package services

import (
	"fmt"

	"github.com/mateusmaaia/showcase-api/domains"
	"github.com/mateusmaaia/showcase-api/repositories"
)

type RealEstateService struct {
	RealEstateRepository repositories.RealEstateRepository
}

func (r *RealEstateService) Insert(realEstate domains.RealEstate) {
	err := realEstate.Validate()

	if err != nil {
		return
	}
	fmt.Printf("RealEstate [%v]", realEstate.ID)
	storeNames := realEstate.DefineStoreNames()

	for _, store := range storeNames {
		r.RealEstateRepository.Insert(store, realEstate)
	}
}

func (r *RealEstateService) FindByStore(storeName string, pageSize int, pageNumber int) ([]domains.RealEstate, int) {
	return r.RealEstateRepository.FindByStore(storeName, pageSize, pageNumber)
}

func (r *RealEstateService) CountByStore(storeName string) int {
	return r.RealEstateRepository.CountByStore(storeName)
}
