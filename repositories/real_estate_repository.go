package repositories

import (
	"fmt"
	"math"

	"github.com/mateusmaaia/showcase-api/domains"
)

type RealEstateRepository struct {
	StoreRealEstates map[string][]domains.RealEstate
}

func (r *RealEstateRepository) Insert(store string, realEstate domains.RealEstate) {
	fmt.Println(fmt.Sprintf("recording store %v", realEstate.ID))
	realEstates := r.StoreRealEstates[store]
	r.StoreRealEstates[store] = append(realEstates, realEstate)
}

func (r *RealEstateRepository) FindByStore(store string, pageSize int, pageNumber int) ([]domains.RealEstate, int) {
	realEstates := r.StoreRealEstates[store]
	total := r.CountByStore(store)
	start := (pageNumber-1)*pageSize + 1

	end := math.Min(float64(start+pageSize), float64(total))

	if start > total {
		return realEstates[0:0], 0
	}

	realEstates = realEstates[start-1 : int(end)]
	return realEstates, len(realEstates)

}

func (r *RealEstateRepository) CountByStore(store string) int {
	return len(r.StoreRealEstates[store])
}
