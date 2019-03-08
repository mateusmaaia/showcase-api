package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNextToZapMinTrue(t *testing.T) {
	validLocation := Location{Lat: zapGroupMinLat, Lon: zapGroupMinLon}
	assert.True(t, validLocation.ItsNextToZapGroup())
}

func TestIsNextToZapMaxTrue(t *testing.T) {
	validLocation := Location{Lat: zapGroupMaxLat, Lon: zapGroupMaxLon}
	assert.True(t, validLocation.ItsNextToZapGroup())
}

func TestIsNextToZapMiddleTrue(t *testing.T) {
	validLocation := Location{Lat: zapGroupMinLat, Lon: zapGroupMaxLon}
	assert.True(t, validLocation.ItsNextToZapGroup())
}

func TestIsNextToZapLonBiggerThanMax(t *testing.T) {
	invalidLocation := Location{Lat: zapGroupMaxLat, Lon: zapGroupMaxLon + 0.1}
	assert.False(t, invalidLocation.ItsNextToZapGroup())
}

func TestIsNextToZapLatBiggerThanMax(t *testing.T) {
	invalidLocation := Location{Lat: zapGroupMaxLat + 0.1, Lon: zapGroupMaxLon}
	assert.False(t, invalidLocation.ItsNextToZapGroup())
}

func TestIsNextToZapLatLowerThanMin(t *testing.T) {
	invalidLocation := Location{Lat: zapGroupMaxLat - 0.1, Lon: zapGroupMaxLon}
	assert.False(t, invalidLocation.ItsNextToZapGroup())
}

func TestIsNextToZapLonLowerThanMin(t *testing.T) {
	invalidLocation := Location{Lat: zapGroupMaxLat, Lon: zapGroupMaxLon - 0.1}
	assert.False(t, invalidLocation.ItsNextToZapGroup())
}
