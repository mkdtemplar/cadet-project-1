package controllers

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/responses"
	"net/http"
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

	}
}

func (sp *ShipController) GetUserPrefPortsName() (*models.UserPreferences, error) {
	id, err := helper.GetQueryID(sp.Request)
	if err != nil {
		return nil, err
	}

	userPreferences, err := sp.IUserPreferencesRepository.FindUserPreferences(sp.Request.Context(), id)
	if err != nil {
		return nil, err
	}

	userPrefPorts, err := sp.IShipPortsRepository.FindUserPrefPorts(sp.Request.Context(), userPreferences)

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

	user, err := sp.IUserRepository.GetById(sp.Request.Context(), id)

	if err != nil {
		return nil, err
	}
	userPorts, err := sp.IShipPortsRepository.FindUserPorts(sp.Request.Context(), user.ID)

	if err != nil {
		return nil, err
	}
	return userPorts, nil
}
