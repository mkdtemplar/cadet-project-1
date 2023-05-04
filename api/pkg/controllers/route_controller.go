package controllers

import (
	"cadet-project/google_API"
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/responses"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"googlemaps.github.io/maps"
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

func (r *RouteController) GetDirections() ([]maps.Route, error) {
	start := helper.GetStartLocation(r.Request)
	end := helper.GetEndLocation(r.Request)
	var err error
	var clientRequest google_API.Request

	clientRequest.Origin, err = r.IShipPortsRepository.GetCityByName(r.Request.Context(), start)

	origin := checkCity(clientRequest.Origin, start)
	if origin == false {
		return nil, errors.New("point of origin do not exist in database")
	}

	clientRequest.Destination, err = r.IShipPortsRepository.GetCityByName(r.Request.Context(), end)
	destination := checkCity(clientRequest.Destination, end)
	if destination == false {
		return nil, errors.New("destination do not exist in database")
	}

	client := google_API.NewClientRequest(clientRequest.Origin, clientRequest.Destination)
	vehicles, err := r.IUserVehicleRepository.FindVehiclesForUser(r.Request.Context(), UserID)
	mileage := maxMileage(vehicles)
	route, err := client.FindRoute(clientRequest, float64(mileage))
	if err != nil {
		return nil, err
	}

	responses.JSON(r.Writer, 200, fmt.Sprintf("%s%d", "Total stops: ", google_API.GetStops()))
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

func checkCity(clientRequest string, start string) bool {
	var err error

	start = strings.Title(strings.ToLower(start))
	if err != nil || clientRequest == "" || clientRequest != start {
		return false
	}

	return true
}
