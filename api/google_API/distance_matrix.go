package google_API

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"cadet-project/pkg/config"
)

func DistanceMatrix(startLat float32, startLong float32, endLat float32, endLong float32) (DistanceMatrixObject, error) {
	distanceMatrix := DistanceMatrixObject{}
	url := fmt.Sprintf("%s, %f, %s, %f, %s, %f, %s, %f, %s, %s",
		"https://maps.googleapis.com/maps/api/distancematrix/json?origins=", startLat, ",", startLong,
		"&destinations=", endLat, ",", endLong, "&key=", config.Config.MapsKey)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return DistanceMatrixObject{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return DistanceMatrixObject{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return DistanceMatrixObject{}, err
	}

	err = json.Unmarshal(body, &distanceMatrix)
	if err != nil {
		log.Println("Cannot unmarshal response")
	}

	return distanceMatrix, err
}
