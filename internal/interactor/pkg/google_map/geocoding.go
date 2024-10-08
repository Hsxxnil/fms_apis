package google_map

import (
	"context"
	"fms/config"
	"fms/internal/interactor/pkg/util/log"
	"googlemaps.github.io/maps"
)

var client *maps.Client

func init() {
	var err error
	client, err = maps.NewClient(maps.WithAPIKey(config.GoogleMapKey))
	if err != nil {
		panic(err)
	}
}

// LatLngToAddress is get address by lat and lng from google-map api
func LatLngToAddress(lat, lng float64) (string, error) {
	if lat == 0 || lng == 0 {
		return "", nil
	}

	reverseResult, err := client.ReverseGeocode(context.Background(), &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lng,
		},
		Language: "zh-TW",
	})
	if err != nil {
		log.Error(err)
		return "", err
	}

	return reverseResult[0].FormattedAddress, nil
}
