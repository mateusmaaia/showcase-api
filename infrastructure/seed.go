package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/mateusmaaia/showcase-api/domains"
	"github.com/mateusmaaia/showcase-api/infrastructure/exceptions"
)

type Seed struct {
	SourceURL string
}

func (s *Seed) Import() ([]domains.RealEstate, error) {
	realEstates := new([]domains.RealEstate)
	resp, err := http.Get(s.SourceURL)
	if err != nil {
		return nil, &exceptions.InvalidUrlError{}
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(realEstates)
	return *realEstates, nil
}
