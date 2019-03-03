package repositories

import "github.com/mateusmaaia/showcase-api/domains"

type RealEstateRepository struct {
	StoreRealEstates map[string][]domains.RealEstate
}

func (r *RealEstateRepository) Insert(store string, realEstate domains.RealEstate) {
	realEstates := r.StoreRealEstates[store]
	r.StoreRealEstates[store] = append(realEstates, realEstate)
}

func (r *RealEstateRepository) FindByStore(store string, pageSize int, offset int) []domains.RealEstate {
	realEstates := r.StoreRealEstates[store]
	return realEstates[offset: pageSize * offset]
}

func (r *RealEstateRepository) CountByStore(store string) int {
	return len(r.StoreRealEstates[store])
}
