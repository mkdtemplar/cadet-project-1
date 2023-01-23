package controllers

import (
	interfaces2 "cadet-project/pkg/interfaces"
	"cadet-project/pkg/responses"
	"net/http"

	"github.com/google/uuid"
)

func NewShipPortsController(IUserRepository interfaces2.IUserRepository, IUserPreferencesRepository interfaces2.IUserPreferencesRepository, IShipPortsRepository interfaces2.IShipPortsRepository) *Controller {
	return &Controller{IUserRepository: IUserRepository, IUserPreferencesRepository: IUserPreferencesRepository, IShipPortsRepository: IShipPortsRepository}
}

func (c *Controller) GetUserPrefPorts(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

	userPreferences, err := c.IUserPreferencesRepository.FindUserPreferences(r.Context(), id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	userPrefPorts, err := c.IShipPortsRepository.FindUserPrefPorts(r.Context(), userPreferences)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userPrefPorts)
}

func (c *Controller) GetUserPorts(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	user, err := c.IUserRepository.GetById(r.Context(), id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	userPorts, err := c.IShipPortsRepository.FindUserPorts(r.Context(), user.ID)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userPorts)
}
