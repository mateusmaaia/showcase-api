package domains

import (
	"strconv"
	"time"
)

type RealEstate struct {
	UsableAreas   int          `json:"usableAreas,omitempty"`
	ListingType   string       `json:"listingType,omitempty"`
	CreatedAt     time.Time    `json:"createdAt,omitempty"`
	ListingStatus string       `json:"listingStatus,omitempty"`
	ID            string       `json:"id,omitempty"`
	ParkingSpaces int          `json:"parkingSpaces,omitempty"`
	UpdatedAt     time.Time    `json:"updatedAt,omitempty"`
	Owner         bool         `json:"owner,omitempty"`
	Images        []string     `json:"images,omitempty"`
	Bathrooms     int          `json:"bathrooms,omitempty"`
	Address       Address      `json:"address,omitempty"`
	Bedrooms      int          `json:"bedrooms,omitempty"`
	PricingInfos  PricingInfos `json:"pricingInfos,omitempty"`
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

	var zapMinRentalPrice float64 = 3500
	var zapMinSalePrice float64 = 600000
	var zapMinSquareMeterPrice float64 = 3500

	var vivaRealMaxRentalPrice float64 = 4000
	var vivaRealMaxSalePrice float64 = 700000

	if r.Address.GeoLocation.Location.ItsNextToZapGroup() {
		zapMinSalePrice = zapMinSalePrice * 0.9
		vivaRealMaxRentalPrice = vivaRealMaxRentalPrice * 1.5
		zapMinSquareMeterPrice = zapMinSquareMeterPrice * 0.9
	}

	if r.PricingInfos.BusinessType == "RENTAL" {

		if rentalTotalPrice >= zapMinRentalPrice {
			storeNames = append(storeNames, StoreName.Zap)
		}

		if rentalTotalPrice <= vivaRealMaxRentalPrice {
			monthlyCondomFee, err := strconv.Atoi(r.PricingInfos.MonthlyCondoFee)
			var maxMonthlyCondomFee float64 = price * 0.3

			if err != nil || float64(monthlyCondomFee) < maxMonthlyCondomFee {
				storeNames = append(storeNames, StoreName.VivaReal)
			}
		}

	}

	if r.PricingInfos.BusinessType == "SALE" {
		if price >= zapMinSalePrice {
			var usableAreas float64 = float64(r.UsableAreas)

			if usableAreas > 0 {
				var meterPrice float64 = price / usableAreas

				if meterPrice > zapMinSquareMeterPrice {
					storeNames = append(storeNames, StoreName.Zap)
				}
			}

			if usableAreas == 0 {
				storeNames = append(storeNames, StoreName.Zap)
			}
		}

		if price <= vivaRealMaxSalePrice {
			storeNames = append(storeNames, StoreName.VivaReal)
		}
	}

	return storeNames
}
