package controllers

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/repository"
	"cadet-project/pkg/repository/generate_id"
	"cadet-project/pkg/repository/validation"
	"cadet-project/pkg/responses"
	"context"
	"net/http"
)

func NewUserPrefController(IUserPreferencesRepository interfaces.IUserPreferencesRepository) *Controller {
	return &Controller{IUserPreferencesRepository: IUserPreferencesRepository}
}
func (c *Controller) ServeHTTPUserPreferences(w http.ResponseWriter, r *http.Request) {
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
	case config.Config.UserPref:
		c.ServeUserPrefEndPoints(w, r)
	}
}

func (c *Controller) CreateUserPref() (*models.UserPreferences, error) {
	v := validation.Validation{}
	userPref, err := helper.ParseUserPrefRequestBody(c.Request)

	validateUserPefData := v.ValidateUserPrefCountry(userPref.UserCountry).ValidateUserId(userPref.UserId)
	if validateUserPefData.Err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, validateUserPefData.Err)
		return nil, validateUserPefData.Err
	}

	userPreferencesStore := repository.NewUserPrefObject(generate_id.GenerateID(), userPref.UserCountry, userPref.UserId)

	err, _ = c.IUserPreferencesRepository.SaveUserPreferences(context.Background(), &userPreferencesStore)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return &models.UserPreferences{}, err
	}

	responses.JSON(c.Writer, http.StatusCreated, userPreferencesStore)

	return &userPreferencesStore, nil
}

func (c *Controller) GetUserPrefById() (*models.UserPreferences, error) {
	id, err := helper.GetQueryID(c.Request)
	if err != nil {
		return nil, err
	}
	return c.IUserPreferencesRepository.FindUserPreferences(context.Background(), id)
}

func (c *Controller) UpdateUserPref() (*models.UserPreferences, error) {
	id, err := helper.GetQueryID(c.Request)
	if err != nil {
		return nil, err
	}

	userPrefFind := &models.UserPreferences{}
	userPrefFind, err = c.IUserPreferencesRepository.FindUserPreferences(context.Background(), id)

	if err != nil {
		return nil, err
	}

	userPrefUpdate, err := helper.ParseUserPrefRequestBody(c.Request)

	validateCountry := V.ValidateUserPrefCountry(userPrefUpdate.UserCountry)
	if validateCountry.Err != nil {
		return nil, validateCountry.Err
	}

	userPrefFind, err = c.IUserPreferencesRepository.UpdateUserPref(context.Background(), id, userPrefUpdate.UserCountry)

	if err != nil {

		return nil, err
	}

	userPrefFind, err = c.IUserPreferencesRepository.FindUserPreferences(context.Background(), id)
	return userPrefFind, nil
}

func (c *Controller) DeleteUserPreferences() error {
	id, err := helper.GetQueryID(c.Request)
	if err != nil {
		return err
	}

	if _, err := c.IUserPreferencesRepository.DeleteUserPreferences(context.Background(), id); err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return err
	}
	responses.JSON(c.Writer, http.StatusNoContent, "")

	return err
}
