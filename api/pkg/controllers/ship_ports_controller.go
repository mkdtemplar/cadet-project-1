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

func NewShipPortsController(IUserRepository interfaces.IUserRepository, IUserPreferencesRepository interfaces.IUserPreferencesRepository, IShipPortsRepository interfaces.IShipPortsRepository) *Controller {
	return &Controller{IUserRepository: IUserRepository, IUserPreferencesRepository: IUserPreferencesRepository, IShipPortsRepository: IShipPortsRepository}
}

func (c *Controller) ServeHTTPShipPorts(w http.ResponseWriter, r *http.Request) {
	c.Writer = w
	c.Request = r

	config.InitConfig("configurations")

	w.Header().Set("content-type", "application/json")

	currentPath := r.URL.Path

	var err error
	var val interface{}

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 401)
		} else {
			responses.JSON(w, http.StatusOK, val)
		}
	}()

	switch currentPath {
	case config.Config.UserPorts:
		val, err = c.GetUserPortsName()
		return

	case config.Config.UserPrefPorts:
		val, err = c.GetUserPrefPortsName()
		return

	case config.Config.PortName:
		val, err = c.GetDirections()
		return
	}
}

func (c *Controller) GetUserPrefPortsName() (*models.UserPreferences, error) {
	id, err := helper.GetQueryID(c.Request)
	if err != nil {
		return nil, err
	}

	userPreferences, err := c.IUserPreferencesRepository.FindUserPreferences(context.Background(), id)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return nil, err
	}

	userPrefPorts, err := c.IShipPortsRepository.FindUserPrefPorts(context.Background(), userPreferences)

	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return nil, err
	}
	responses.JSON(c.Writer, http.StatusOK, userPrefPorts)
	return userPrefPorts, nil
}

func (c *Controller) GetUserPortsName() (*models.User, error) {
	id, err := helper.GetQueryID(c.Request)
	if err != nil {
		return nil, err
	}

	user, err := c.IUserRepository.GetById(context.Background(), id)

	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return nil, err
	}

	userPorts, err := c.IShipPortsRepository.FindUserPorts(context.Background(), user.ID)

	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return nil, err
	}
	responses.JSON(c.Writer, http.StatusOK, userPorts)
	return userPorts, nil
}

func (c *Controller) GetDirections() ([]maps.Route, error) {
	start := helper.GetQueryStart(c.Request)
	end := helper.GetQueryEnd(c.Request)
	var err error
	var clientRequest google_API.ClientData

	clientRequest.Origin, err = c.IShipPortsRepository.GetCityByName(context.Background(), start)
	if err != nil || clientRequest.Origin == "" || clientRequest.Origin != start {
		return nil, errors.New("point of origin do not exist in database")
	}

	clientRequest.Destination, err = c.IShipPortsRepository.GetCityByName(context.Background(), end)
	if err != nil || clientRequest.Destination == "" || clientRequest.Destination != end {
		return nil, errors.New("destination do not exist in database")
	}

	route := google_API.NewClientData(clientRequest.Origin, clientRequest.Destination)

	return route.FindRoute(), nil
}
