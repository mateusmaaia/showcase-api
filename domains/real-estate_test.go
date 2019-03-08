package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var locationOutOfBoundingBox = Address{GeoLocation: GeoLocation{Location: Location{Lat: 0, Lon: -47.693419}}}

var locationInBoundingBox = Address{GeoLocation: GeoLocation{Location: Location{Lat: -23.568704, Lon: -46.693419}}}

// Location
func TestValidateGeoLocationWithLonAndLatEqualZero(t *testing.T) {
	invalidLocation := Location{Lat: 0, Lon: 0}
	realEstate := RealEstate{Address: Address{GeoLocation: GeoLocation{Location: invalidLocation}}}
	assert.Error(t, realEstate.Validate())
}

func TestValidateGeoLocationWithLonAndLatNotZero(t *testing.T) {
	invalidLocation := Location{Lat: -23.502555, Lon: -46.716542}
	realEstate := RealEstate{Address: Address{GeoLocation: GeoLocation{Location: invalidLocation}}}
	assert.Nil(t, realEstate.Validate())
}

func TestValidateGeoLocationWithOnlyLonEqualZero(t *testing.T) {
	invalidLocation := Location{Lat: -23.502555, Lon: 0}
	realEstate := RealEstate{Address: Address{GeoLocation: GeoLocation{Location: invalidLocation}}}
	assert.Nil(t, realEstate.Validate())
}

func TestValidateGeoLocationWithOnlyLatEqualZero(t *testing.T) {
	invalidLocation := Location{Lat: 0, Lon: -46.716542}
	realEstate := RealEstate{Address: Address{GeoLocation: GeoLocation{Location: invalidLocation}}}
	assert.Nil(t, realEstate.Validate())
}

// Rental
func TestValidateStoreRentalTotalPriceOutOfBoudingBoxZapAndVivaReal(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "RENTAL", RentalTotalPrice: "3500"}
	realEstate := RealEstate{PricingInfos: pricingInfos, Address: locationOutOfBoundingBox}
	assert.Contains(t, realEstate.DefineStoreNames(), "viva-real")
	assert.Contains(t, realEstate.DefineStoreNames(), "zap")
}

func TestValidateStoreRentalTotalPriceOutOfBoudingBoxOnlyZap(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "RENTAL", RentalTotalPrice: "6000"}
	realEstate := RealEstate{PricingInfos: pricingInfos, Address: locationOutOfBoundingBox}
	assert.NotContains(t, realEstate.DefineStoreNames(), "viva-real")
	assert.Contains(t, realEstate.DefineStoreNames(), "zap")
}

func TestValidateStoreRentalTotalPriceOutOfBoudingBoxOnlyVivaReal(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "RENTAL", RentalTotalPrice: "3499"}
	realEstate := RealEstate{PricingInfos: pricingInfos, Address: locationOutOfBoundingBox}
	assert.NotContains(t, realEstate.DefineStoreNames(), "zap")
	assert.Contains(t, realEstate.DefineStoreNames(), "viva-real")
}

func TestValidateStoreRentalTotalPrice50PercentBiggerInBoudingBoxZapAndVivaReal(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "RENTAL", RentalTotalPrice: "6000"}
	realEstate := RealEstate{PricingInfos: pricingInfos, Address: locationInBoundingBox}
	assert.Contains(t, realEstate.DefineStoreNames(), "viva-real")
	assert.Contains(t, realEstate.DefineStoreNames(), "zap")
}

func TestValidateStoreRentalTotalPrice50PercentBiggerInBoudingBoxOnlyZap(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "RENTAL", RentalTotalPrice: "6001"}
	realEstate := RealEstate{PricingInfos: pricingInfos, Address: locationInBoundingBox}
	assert.NotContains(t, realEstate.DefineStoreNames(), "viva-real")
	assert.Contains(t, realEstate.DefineStoreNames(), "zap")
}

func TestValidateStoreMonthlyCondoFeeBiggerThen30PercentInBoudingBox(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "RENTAL", RentalTotalPrice: "6000", MonthlyCondoFee: "1900", Price: "4100"}
	realEstate := RealEstate{PricingInfos: pricingInfos, Address: locationInBoundingBox}
	assert.NotContains(t, realEstate.DefineStoreNames(), "viva-real")
}

func TestValidateStoreMonthlyCondoFeeLowerThen30PercentInBoudingBox(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "RENTAL", RentalTotalPrice: "6000", MonthlyCondoFee: "500", Price: "5500"}
	realEstate := RealEstate{PricingInfos: pricingInfos, Address: locationInBoundingBox}
	assert.Contains(t, realEstate.DefineStoreNames(), "viva-real")
}

func TestValidateStoreMonthlyCondoFeeLowerThen30PercentOutOfBoudingBox(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "RENTAL", RentalTotalPrice: "3612", MonthlyCondoFee: "812", Price: "2800"}
	realEstate := RealEstate{PricingInfos: pricingInfos, Address: locationInBoundingBox}
	assert.Contains(t, realEstate.DefineStoreNames(), "viva-real")
}

func TestValidateStoreMonthlyCondoBiggerThen30PercentOutOfBoudingBox(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "RENTAL", RentalTotalPrice: "4000", MonthlyCondoFee: "1200", Price: "2800"}
	realEstate := RealEstate{PricingInfos: pricingInfos, Address: locationInBoundingBox}
	assert.NotContains(t, realEstate.DefineStoreNames(), "viva-real")
}

func TestValidateStoreMonthlyCondoNotANumberOutOfBoudingBox(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "RENTAL", RentalTotalPrice: "3100", MonthlyCondoFee: "abc"}
	realEstate := RealEstate{PricingInfos: pricingInfos, Address: locationInBoundingBox}
	assert.Contains(t, realEstate.DefineStoreNames(), "viva-real")
}

// Sale
func TestValidateStoreSalePriceOutOfBoudingBoxZapAndVivaReal(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "SALE", Price: "600000"}
	realEstate := RealEstate{UsableAreas: 80, PricingInfos: pricingInfos, Address: locationOutOfBoundingBox}
	assert.Contains(t, realEstate.DefineStoreNames(), "viva-real")
	assert.Contains(t, realEstate.DefineStoreNames(), "zap")
}

func TestValidateStoreSalePriceOutOfBoudingBoxWithNoValidUsableAreasZapAndVivaReal(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "SALE", Price: "600000"}
	realEstate := RealEstate{UsableAreas: 0, PricingInfos: pricingInfos, Address: locationOutOfBoundingBox}
	assert.Contains(t, realEstate.DefineStoreNames(), "viva-real")
	assert.Contains(t, realEstate.DefineStoreNames(), "zap")
}

func TestValidateStoreSalePriceOutOfBoudingBoxWithUsableAreasLowerThanMinimumOnlyVivaReal(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "SALE", Price: "600000"}
	realEstate := RealEstate{UsableAreas: 1200, PricingInfos: pricingInfos, Address: locationOutOfBoundingBox}
	assert.Contains(t, realEstate.DefineStoreNames(), "viva-real")
	assert.NotContains(t, realEstate.DefineStoreNames(), "zap")
}

func TestValidateStoreSalePriceInBoudingBoxWithUsableAreasBiggerThanMinimumOnlyZap(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "SALE", Price: "900000"}
	realEstate := RealEstate{UsableAreas: 200, PricingInfos: pricingInfos, Address: locationInBoundingBox}
	assert.NotContains(t, realEstate.DefineStoreNames(), "viva-real")
	assert.Contains(t, realEstate.DefineStoreNames(), "zap")
}

func TestValidateStoreSalePriceInBoudingBoxWithUsableAreasLowerThanMinimum(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "SALE", Price: "900000"}
	realEstate := RealEstate{UsableAreas: 290, PricingInfos: pricingInfos, Address: locationInBoundingBox}
	assert.NotContains(t, realEstate.DefineStoreNames(), "viva-real")
	assert.NotContains(t, realEstate.DefineStoreNames(), "zap")
}

func TestValidateStoreSalePriceInBoudingBoxOnlyVivaReal(t *testing.T) {
	pricingInfos := PricingInfos{BusinessType: "SALE", Price: "400000"}
	realEstate := RealEstate{UsableAreas: 20, PricingInfos: pricingInfos, Address: locationInBoundingBox}
	assert.Contains(t, realEstate.DefineStoreNames(), "viva-real")
	assert.NotContains(t, realEstate.DefineStoreNames(), "zap")
}
