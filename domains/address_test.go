package domains

import (
	"testing"
)

func TestIsNextToZap(t *testing.T) {
	tests := []struct{
		name string
		lat float64
		lon float64
		expected bool
	}{
		{"Min_LatAndLon_True", zapGroupMinLat, zapGroupMinLon,  true},
		{"Max_LatAndLon_True", zapGroupMaxLat, zapGroupMaxLon,true},
		{"Middle_LatAndLon_True", zapGroupMinLat, zapGroupMaxLon,true},
		{"Lon_BiggerThanMax_False", zapGroupMaxLat, zapGroupMaxLon + 0.1,false},
		{"Lat_BiggerThanMax_False", zapGroupMaxLat + 0.1, zapGroupMaxLon,false},
		{"Lat_LowerThanMin_False", zapGroupMaxLat - 0.1, zapGroupMaxLon,false},
		{"Lon_LowerThanMin_False", zapGroupMaxLat, zapGroupMaxLon - 0.1,false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validLocation := Location{Lat: tt.lat, Lon: tt.lon}

			got := validLocation.ItsNextToZapGroup()

			if got != tt.expected {
				t.Errorf("Type of realEstate.Validate() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
