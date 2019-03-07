package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestValidateStoreNameZap(t *testing.T) {

}
