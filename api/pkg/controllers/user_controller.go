package controllers

import (
	"cadet-project/pkg/controllers/helper"
	interfaces2 "cadet-project/pkg/interfaces"
	"cadet-project/pkg/repository/validation"
	"cadet-project/pkg/responses"
	"net/http"

	"github.com/google/uuid"
)

func NewUserController(IUserRepository interfaces2.IUserRepository, IUserPreferencesRepository interfaces2.IUserPreferencesRepository,
	IShipPortsRepository interfaces2.IShipPortsRepository) *Controller {
	return &Controller{IUserRepository: IUserRepository, IUserPreferencesRepository: IUserPreferencesRepository, IShipPortsRepository: IShipPortsRepository}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	v := validation.Validation{}
	user := helper.ParseUserRequestBody(w, r)

	checkCredentials := v.ValidateUserEmail(user.Email).ValidateUserName(user.Name)

	if checkCredentials.Err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, checkCredentials.Err)
		return
	}

	if _, err := c.IUserRepository.Create(r.Context(), user); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, user)

}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

	if _, err := c.IUserRepository.Delete(r.Context(), id); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}

func (c *Controller) GetId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	user, err := c.IUserRepository.GetById(r.Context(), id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}
