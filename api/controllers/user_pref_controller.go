package controllers

import (
	"cadet-project/controllers/helper"
	"cadet-project/interfaces"
	"cadet-project/models"
	"cadet-project/repository"
	"cadet-project/repository/generate_id"
	"cadet-project/repository/validation"
	"cadet-project/responses"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func NewUserPrefController(IUserPreferencesRepository interfaces.IUserPreferencesRepository) *Controller {
	return &Controller{IUserPreferencesRepository: IUserPreferencesRepository}
}

func (c *Controller) CreateUserPreferences(w http.ResponseWriter, r *http.Request) {
	v := validation.Validation{}
	userPref := helper.ParseUserPrefRequestBody(w, r)

	validateUserPefData := v.ValidateUserPrefCountry(userPref.UserCountry).ValidateUserId(userPref.UserId)
	if validateUserPefData.Err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, validateUserPefData.Err)
		return
	}

	userPreferencesStore := repository.NewUserPrefObject(generate_id.GenerateID(), userPref.UserCountry, userPref.UserId)

	_, err := c.IUserPreferencesRepository.SaveUserPreferences(r.Context(), &userPreferencesStore)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, userPreferencesStore)
}

func (c *Controller) GetUserPreference(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

	userPreferences, err := c.IUserPreferencesRepository.FindUserPreferences(r.Context(), id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userPreferences)
}

func (c *Controller) UpdateUserPreferences(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var err error
	v := validation.Validation{}
	userPrefFind := &models.UserPreferences{}
	userPrefFind, err = c.IUserPreferencesRepository.FindUserPreferences(r.Context(), id)

	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("user preferences not found"))
	}

	userPrefUpdate := helper.ParseUserPrefRequestBody(w, r)

	validateCountry := v.ValidateUserPrefCountry(userPrefUpdate.UserCountry)
	if validateCountry.Err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, validateCountry.Err)
		return
	}

	userPrefFind, err = c.IUserPreferencesRepository.UpdateUserPref(r.Context(), id, userPrefUpdate.UserCountry)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	userPrefFind, err = c.IUserPreferencesRepository.FindUserPreferences(r.Context(), id)
	responses.JSON(w, http.StatusOK, userPrefFind)
}

func (c *Controller) DeleteUserPref(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

	if _, err := c.IUserPreferencesRepository.DeleteUserPreferences(r.Context(), id); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}
