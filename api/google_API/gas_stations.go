package google_API

import (
	"cadet-project/pkg/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GasStations(latitude float64, longitude float64, radius int) (GasStationsObject, error) {
	url := fmt.Sprintf("%s%f%s%f%s%d%s%s%s", "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=",
		latitude, ",%20", longitude, "&radius=", radius, "&type=gas_station&", "key=", config.Config.MapsKey)

	resp, err := http.Get(url)
	if err != nil {
		return GasStationsObject{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GasStationsObject{}, err
	}

	mapResponse := GasStationsObject{}

	err = json.Unmarshal(body, &mapResponse)
	for _, i := range mapResponse.Results {
		fmt.Println(i.BusinessStatus)
		fmt.Println(i.Name)
		fmt.Println(i.Geometry.Location)
		fmt.Println(i.OpeningHours.OpenNow)
		fmt.Println(i.Vicinity)
	}

	return mapResponse, nil
}
