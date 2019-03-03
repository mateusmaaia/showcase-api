package domains

import (
	"strconv"
	"time"
)

type RealEstate struct {
	UsableAreas   int          `json:"usableAreas"`
	ListingType   string       `json:"listingType"`
	CreatedAt     time.Time    `json:"createdAt"`
	ListingStatus string       `json:"listingStatus"`
	ID            string       `json:"id"`
	ParkingSpaces int          `json:"parkingSpaces"`
	UpdatedAt     time.Time    `json:"updatedAt"`
	Owner         bool         `json:"owner"`
	Images        []string     `json:"images"`
	Bathrooms     int          `json:"bathrooms"`
	Address       Address      `json:"address"`
	Bedrooms      int          `json:"bedrooms"`
	PricingInfos  PricingInfos `json:"pricingInfos"`
}

type store = string

type list struct {
	Zap      store
	VivaReal store
}

var StoreName = &list{
	Zap:      "zap",
	VivaReal: "viva-real",
}

func (r *RealEstate) Validate() error {
	var err error

	err = r.Address.GeoLocation.Location.IsLocationValid()
	if err != nil {
		return err
	}

	return nil
}

func (r *RealEstate) DefineStoreNames() []string {
	storeNames := make([]string, 0)
	rentalTotalPrice, _ := strconv.ParseFloat(r.PricingInfos.RentalTotalPrice, 64)
	price, _ := strconv.ParseFloat(r.PricingInfos.Price, 64)

	var zapMaxRentalPrice float64 = 3500
	var zapMinSalePrice float64 = 600000

	var vivaRealMaxRentalPrice float64 = 4000
	var vivaRealMaxSalePrice float64 = 700000

	if r.Address.GeoLocation.Location.ItsNextToZapGroup() {
		zapMinSalePrice = zapMinSalePrice * 0.9
		vivaRealMaxRentalPrice = vivaRealMaxRentalPrice * 1.5
	}

	if r.PricingInfos.BusinessType == "RENTAL" {

		if rentalTotalPrice > zapMaxRentalPrice {
			storeNames = append(storeNames, StoreName.Zap)
		}

		if rentalTotalPrice < vivaRealMaxRentalPrice {
			monthlyCondomFee, err := strconv.Atoi(r.PricingInfos.MonthlyCondoFee)
			var maxMonthlyCondomFee float64 = rentalTotalPrice * 0.3

			if err != nil || monthlyCondomFee > 0 && float64(monthlyCondomFee) < maxMonthlyCondomFee {
				storeNames = append(storeNames, StoreName.VivaReal)
			}
		}

	}

	if r.PricingInfos.BusinessType == "SALE" {
		if price >= zapMinSalePrice {
			if r.UsableAreas == 0 || r.UsableAreas > 3500 {
				storeNames = append(storeNames, StoreName.Zap)
			}
		}

		if price <= vivaRealMaxSalePrice {
			storeNames = append(storeNames, StoreName.VivaReal)
		}
	}

	return storeNames
}
