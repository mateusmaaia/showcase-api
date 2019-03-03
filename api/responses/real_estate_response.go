package responses

import "github.com/mateusmaaia/showcase-api/domains"

type RealEstateResponse struct {
	PageNumber int                  `json:"pageNumber"`
	PageSize   int                  `json:"pageSize"`
	TotalCount int                  `json:"totalCount"`
	Listings   []domains.RealEstate `json:"listings"`
}
