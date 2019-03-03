package domains

import "github.com/mateusmaaia/showcase-api/domains/exceptions"

const (
	zapGroupMinLon = -46.693419
	zapGroupMinLat = -23.568704
	zapGroupMaxLon = -46.641146
	zapGroupMaxLat = -23.546686
)

type Location struct {
	Lon int `json:"lon"`
	Lat int `json:"lat"`
}

type GeoLocation struct {
	Precision string   `json:"precision"`
	Location  Location `json:"location"`
}

type Address struct {
	City         string      `json:"city"`
	Neighborhood string      `json:"neighborhood"`
	GeoLocation  GeoLocation `json:"geoLocation"`
}

func (l *Location) IsLocationValid() error {

	if l.Lat == 0 && l.Lon == 0 {
		return &exceptions.InvalidLocationError{}
	}

	return nil
}

func (l *Location) ItsNextToZapGroup() bool {
	var lat float64 = float64(l.Lat)
	var lon float64 = float64(l.Lon)
	return lat >= zapGroupMinLat && lat <= zapGroupMaxLat && lon >= zapGroupMinLon && lon <= zapGroupMaxLon
}
