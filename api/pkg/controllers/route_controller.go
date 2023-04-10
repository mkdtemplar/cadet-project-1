package controllers

import (
	"cadet-project/google_API"
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/responses"
	"errors"
	"net/http"
)

func NewRouteController(IUserVehicleRepository interfaces.IUserVehicleRepository, IShipPortsRepository interfaces.IShipPortsRepository) *RouteController {
	return &RouteController{IUserVehicleRepository: IUserVehicleRepository, IShipPortsRepository: IShipPortsRepository}
}

func (r *RouteController) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	r.Request = rq
	r.Writer = w
	config.InitConfig("configurations")

	w.Header().Set("Content-Type", "application/json")

	currentPath := rq.URL.Path

	var err error
	var val interface{}

	defer func() {
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		} else {
			responses.JSON(w, http.StatusOK, val)
		}
	}()

	switch currentPath {
	case config.Config.PortName:
		val, err = r.GetDirections()
		return
	}
}

func (r *RouteController) GetDirections() ([]google_API.Route, error) {
	start := helper.GetQueryStart(r.Request)
	end := helper.GetQueryEnd(r.Request)
	var err error
	var clientRequest google_API.Request
	var startLat float64
	var startLng float64
	var endLat float64
	var endLng float64

	clientRequest.Origin, err = r.IShipPortsRepository.GetCityByName(r.Request.Context(), start)
	if err != nil || clientRequest.Origin == "" || clientRequest.Origin != start {
		return nil, errors.New("point of origin do not exist in database")
	}

	clientRequest.Destination, err = r.IShipPortsRepository.GetCityByName(r.Request.Context(), end)
	if err != nil || clientRequest.Destination == "" || clientRequest.Destination != end {
		return nil, errors.New("destination do not exist in database")
	}

	originLatitude, err := r.IShipPortsRepository.GetCityLatitude(r.Request.Context(), clientRequest.Origin)
	originLongitude, err := r.IShipPortsRepository.GetCityLongitude(r.Request.Context(), clientRequest.Origin)

	destinationLatitude, err := r.IShipPortsRepository.GetCityLatitude(r.Request.Context(), clientRequest.Destination)
	destinationLongitude, err := r.IShipPortsRepository.GetCityLongitude(r.Request.Context(), clientRequest.Destination)

	client := google_API.NewClientRequest(clientRequest.Origin, clientRequest.Destination)
	route, err := client.FindRoute(clientRequest)
	if err != nil {
		return nil, err
	}

	totalDistance := google_API.GetTotalDistance()

	vehicles, err := r.IUserVehicleRepository.FindVehiclesForUser(r.Request.Context(), UserID)
	mileage := maxMileage(vehicles)

	if len(vehicles) == 1 {
		if float32(totalDistance) <= vehicles[0].Mileage {
			return route, nil
		} else {
			gasStations, _ := google_API.GasStations(originLatitude, originLongitude, totalDistance)
			for i := 0; i < len(gasStations.Results); i++ {
				startLat = gasStations.Results[i].Geometry.Location.Lat
				startLng = gasStations.Results[i].Geometry.Location.Lng
				endLat = gasStations.Results[i+1].Geometry.Location.Lat
				endLng = gasStations.Results[i+1].Geometry.Location.Lng

				distanceGasStations, _ := google_API.DistanceMatrix(float32(startLat), float32(startLng), float32(endLat), float32(endLng))
				for j := 0; j < len(distanceGasStations.Rows); j++ {
					for k := 0; k < len(distanceGasStations.Rows[i].Elements); k++ {
						if distanceGasStations.Rows[i].Elements[k].Distance.Value/1000 > int(mileage) || gasStations.Results[i].OpeningHours.OpenNow == false {
							return nil, errors.New("route not possible vehicle can not reach the gas station for refueling")
						}
					}
				}
			}
			startLat = gasStations.Results[len(gasStations.Results)-1].Geometry.Location.Lat
			startLng = gasStations.Results[len(gasStations.Results)-1].Geometry.Location.Lng
			distanceGasStations, _ := google_API.DistanceMatrix(float32(startLat), float32(startLng), destinationLatitude, destinationLongitude)
			if distanceGasStations.Rows[0].Elements[0].Distance.Value/1000 > int(mileage) || gasStations.Results[0].OpeningHours.OpenNow == false {
				return nil, errors.New("route not possible vehicle can not reach the gas station for refueling")
			}
		}
	}
	return route, nil
}

func maxMileage(vehicles []*models.Vehicle) float32 {
	max := vehicles[0].Mileage

	for _, m := range vehicles {
		if max < m.Mileage {
			max = m.Mileage
		}
	}

	return max
}
