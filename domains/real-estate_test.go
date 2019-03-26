package domains

import (
	"github.com/mateusmaaia/showcase-api/domains/exceptions"
	"reflect"
	"strings"
	"testing"
)


func TestLocation(t *testing.T) {
	tests := []struct{
		name string
		lat float64
		lon float64
		expected error
	}{
		{"GeoLocationWithLonAndLatEqualZero Error", 0, 0,  &exceptions.InvalidLocationError{}},
		{"GeoLocationWithLonAndLatNotZero Nil", -23.502555, -46.716542,nil},
		{"GeoLocationWithOnlyLonEqualZero Nil", -23.502555, 0,nil},
		{"GeoLocationWithOnlyLonEqualZero Nil", 0, -46.716542,nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invalidLocation := Location{Lat: tt.lat, Lon: tt.lon}
			realEstate := RealEstate{Address: Address{GeoLocation: GeoLocation{Location: invalidLocation}}}
			got := realEstate.Validate()

			if reflect.TypeOf(got) != reflect.TypeOf(tt.expected) {
				t.Errorf("Type of realEstate.Validate() = %v, expected %v", reflect.TypeOf(got), reflect.TypeOf(tt.expected))
			}
		})
	}

}

func TestRentalAndSale(t *testing.T){

	locationOutOfBoundingBox := Address{GeoLocation: GeoLocation{Location: Location{Lat: 0, Lon: -47.693419}}}
	locationInBoundingBox := Address{GeoLocation: GeoLocation{Location: Location{Lat: -23.568704, Lon: -46.693419}}}

	tests := []struct {
		name string
		businessType string
		rentalTotalPrice string
		monthlyCondoFee string
		price string
		usableAreas int
		isZapBoudingBox bool
		expected string
	}{
		{"Rental_TotalPrice_OutOfBoudingBox ZapAndVivaReal", "RENTAL", "3500", "", "", 0, false, "zap viva-real"},
		{"Rental TotalPrice OutOfBoudingBox OnlyZap", "RENTAL", "6000", "", "", 0, false,"zap"},
		{"Rental TotalPrice OutOfBoudingBox OnlyVivaReal", "RENTAL", "3499", "", "", 0, false,"viva-real"},
		{"Rental TotalPrice50PercentBigger InBoudingBox ZapAndVivaReal", "RENTAL", "6000", "", "", 0, true,"zap viva-real"},
		{"Rental TotalPrice50PercentBigger InBoudingBox OnlyZap", "RENTAL", "6001", "", "", 0, true,"zap"},
		{"Rental MonthlyCondoFeeBiggerThen30Percent InBoudingBox OnlyZap", "RENTAL", "6000", "1900", "4100", 0, true,"zap"},
		{"Rental MonthlyCondoFeeLowerThen30Percent InBoudingBox ZapAndVivaReal", "RENTAL", "6000", "500", "5500", 0, true,"zap viva-real"},
		{"Rental MonthlyCondoFeeLowerThen30Percent OutBoudingBox ZapAndVivaReal", "RENTAL", "3612", "812", "2800", 0, false,"zap viva-real"},
		{"Rental MonthlyCondoFeeBiggerThen30Percent OutBoudingBox ZapAndVivaReal", "RENTAL", "4000", "1200", "2800", 0, false,"zap"},
		{"Rental MonthlyCondoNotANumber OutOfBoudingBox None", "RENTAL", "3100", "abc", "", 0, false,"viva-real"},
		{"Sale Price OutOfBoudingBox ZapAndVivaReal", "SALE", "", "", "600000", 0,false,"zap viva-real"},
		{"Sale Price OutOfBoudingBox WithNoValidUsableAreas ZapAndVivaReal", "SALE", "", "", "600000", 0,false,"zap viva-real"},
		{"Sale Price OutOfBoudingBox WithUsableAreasLowerThanMinimum OnlyVivaReal", "SALE", "", "", "600000", 1200,false,"viva-real"},
		{"Sale Price InBoudingBox WithUsableAreasBiggerThanMinimum OnlyZap", "SALE", "", "", "900000", 200,true,"zap"},
		{"Sale Price InBoudingBox WithUsableAreasLowerThanMinimum None", "SALE", "", "", "900000", 290,true,""},
		{"Sale Price InBoudingBox OnlyVivaReal", "SALE", "", "", "400000", 20,true,"viva-real"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			pricingInfos := PricingInfos{BusinessType: tt.businessType, RentalTotalPrice: tt.rentalTotalPrice, MonthlyCondoFee: tt.monthlyCondoFee, Price: tt.price}

			location := locationOutOfBoundingBox

			if tt.isZapBoudingBox {
				location = locationInBoundingBox
			}

			realEstate := RealEstate{PricingInfos: pricingInfos, Address:location, UsableAreas: tt.usableAreas}

			got := realEstate.DefineStoreNames()

			if strings.Join(got, " ") != tt.expected {
				t.Errorf("DefineStoreNames() = %v, expected %v", got, tt.expected)
			}
		})
	}
}