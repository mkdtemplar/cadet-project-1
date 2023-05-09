package controllers

import (
	"cadet-project/google_API"
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/responses"
	"cadet-project/pkg/utils"
	"errors"
	"fmt"
	"net/http"

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
	var val1 any

	defer func() {
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		} else {
			responses.JSON(w, http.StatusOK, val)
			responses.JSON(w, http.StatusOK, val1)
		}
	}()

	switch currentPath {
	case config.Config.PortName:
		val, val1, err = r.GetDirections()
		return
	}
}

func (r *RouteController) GetDirections() ([]maps.Route, []google_API.GasStationsObject, error) {
	start := helper.GetStartLocation(r.Request)
	end := helper.GetEndLocation(r.Request)
	var err error
	var clientRequest google_API.Request

	clientRequest.Origin, err = r.IShipPortsRepository.GetCityByName(r.Request.Context(), start)

	origin := utils.CheckCity(clientRequest.Origin, start)
	if origin == false {
		return nil, nil, errors.New("point of origin do not exist in database")
	}

	clientRequest.Destination, err = r.IShipPortsRepository.GetCityByName(r.Request.Context(), end)
	destination := utils.CheckCity(clientRequest.Destination, end)
	if destination == false {
		return nil, nil, errors.New("destination do not exist in database")
	}

	client := google_API.NewClientRequest(clientRequest.Origin, clientRequest.Destination)
	vehicles, err := r.IUserVehicleRepository.FindVehiclesForUser(r.Request.Context(), UserID)
	mileage := utils.MaxMileage(vehicles)
	route, gasStations, err := client.FindRoute(clientRequest, float64(mileage))
	if err != nil {
		return nil, nil, err
	}

	responses.JSON(r.Writer, 200, fmt.Sprintf("%s%d", "Total stops: ", google_API.GetStops()))
	return route, gasStations, nil
}
