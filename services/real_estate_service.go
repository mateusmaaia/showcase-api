package services

import (
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

	storeNames := realEstate.DefineStoreNames()

	for _, store := range storeNames {
		r.RealEstateRepository.Insert(store, realEstate)
	}
}

func (r *RealEstateService) FindByStore(storeName string, pageSize int, offset int) []domains.RealEstate {
	return r.RealEstateRepository.FindByStore(storeName, pageSize, offset)
}

func (r *RealEstateService) CountByStore(storeName string) int {
	return r.RealEstateRepository.CountByStore(storeName)
}
