package controllers

import (
	"cadet-project/google_API"
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/responses"
	"context"
	"errors"
	"net/http"

	"googlemaps.github.io/maps"
)

func NewShipPortsController(IUserRepository interfaces.IUserRepository, IUserPreferencesRepository interfaces.IUserPreferencesRepository, IShipPortsRepository interfaces.IShipPortsRepository) *ShipController {
	return &ShipController{IUserRepository: IUserRepository, IUserPreferencesRepository: IUserPreferencesRepository, IShipPortsRepository: IShipPortsRepository}
}

func (sp *ShipController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sp.Writer = w
	sp.Request = r

	config.InitConfig("configurations")

	w.Header().Set("content-type", "application/json")

	currentPath := r.URL.Path

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
	case config.Config.UserPorts:
		val, err = sp.GetUserPortsName()
		return

	case config.Config.UserPrefPorts:
		val, err = sp.GetUserPrefPortsName()
		return

	case config.Config.PortName:
		val, err = sp.GetDirections()
		return
	}
}

func (sp *ShipController) GetUserPrefPortsName() (*models.UserPreferences, error) {
	id, err := helper.GetQueryID(sp.Request)
	if err != nil {
		return nil, err
	}

	userPreferences, err := sp.IUserPreferencesRepository.FindUserPreferences(context.Background(), id)
	if err != nil {
		return nil, err
	}

	userPrefPorts, err := sp.IShipPortsRepository.FindUserPrefPorts(context.Background(), userPreferences)

	if err != nil {
		return nil, err
	}
	return userPrefPorts, nil
}

func (sp *ShipController) GetUserPortsName() (*models.User, error) {
	id, err := helper.GetQueryID(sp.Request)
	if err != nil {
		return nil, err
	}

	user, err := sp.IUserRepository.GetById(context.Background(), id)

	if err != nil {
		return nil, err
	}
	userPorts, err := sp.IShipPortsRepository.FindUserPorts(context.Background(), user.ID)

	if err != nil {
		return nil, err
	}
	return userPorts, nil
}

func (sp *ShipController) GetDirections() ([]maps.Route, error) {
	start := helper.GetQueryStart(sp.Request)
	end := helper.GetQueryEnd(sp.Request)
	var err error
	var clientRequest google_API.ClientData

	clientRequest.Origin, err = sp.IShipPortsRepository.GetCityByName(context.Background(), start)
	if err != nil || clientRequest.Origin == "" || clientRequest.Origin != start {
		return nil, errors.New("point of origin do not exist in database")
	}

	clientRequest.Destination, err = sp.IShipPortsRepository.GetCityByName(context.Background(), end)
	if err != nil || clientRequest.Destination == "" || clientRequest.Destination != end {
		return nil, errors.New("destination do not exist in database")
	}

	route := google_API.NewClientData(clientRequest.Origin, clientRequest.Destination)

	return route.FindRoute(), nil
}
