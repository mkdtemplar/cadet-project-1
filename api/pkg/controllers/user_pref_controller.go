package controllers

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/repository"
	"cadet-project/pkg/repository/generate_id"
	"cadet-project/pkg/responses"
	"context"
	"net/http"
)

func NewUserPrefController(IUserPreferencesRepository interfaces.IUserPreferencesRepository) *UserPrefController {
	return &UserPrefController{IUserPreferencesRepository: IUserPreferencesRepository}
}
func (upc *UserPrefController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upc.Writer = w
	upc.Request = r

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
		upc.ServeUserPrefEndPoints(w, r)
	}
}

func (upc *UserPrefController) CreateUserPref() (*models.UserPreferences, error) {

	userPref, err := helper.ParseUserPrefRequestBody(upc.Request)

	validateUserPefData := V.ValidateUserPrefCountry(userPref.UserCountry).ValidateUserId(userPref.UserId)
	if validateUserPefData.Err != nil {
		responses.ERROR(upc.Writer, http.StatusUnprocessableEntity, validateUserPefData.Err)
		return nil, validateUserPefData.Err
	}

	userPreferencesStore := repository.NewUserPrefObject(generate_id.GenerateID(), userPref.UserCountry, userPref.UserId)

	err, _ = upc.IUserPreferencesRepository.SaveUserPreferences(context.Background(), &userPreferencesStore)
	if err != nil {
		responses.ERROR(upc.Writer, http.StatusUnprocessableEntity, err)
		return &models.UserPreferences{}, err
	}

	responses.JSON(upc.Writer, http.StatusCreated, userPreferencesStore)

	return &userPreferencesStore, nil
}

func (upc *UserPrefController) GetUserPrefById() (*models.UserPreferences, error) {
	id, err := helper.GetQueryID(upc.Request)
	if err != nil {
		return nil, err
	}
	return upc.IUserPreferencesRepository.FindUserPreferences(context.Background(), id)
}

func (upc *UserPrefController) UpdateUserPref() (*models.UserPreferences, error) {
	id, err := helper.GetQueryID(upc.Request)
	if err != nil {
		return nil, err
	}

	userPrefFind := &models.UserPreferences{}
	userPrefFind, err = upc.IUserPreferencesRepository.FindUserPreferences(context.Background(), id)

	if err != nil {
		return nil, err
	}

	userPrefUpdate, err := helper.ParseUserPrefRequestBody(upc.Request)

	validateCountry := V.ValidateUserPrefCountry(userPrefUpdate.UserCountry)
	if validateCountry.Err != nil {
		return nil, validateCountry.Err
	}

	userPrefFind, err = upc.IUserPreferencesRepository.UpdateUserPref(context.Background(), id, userPrefUpdate.UserCountry)

	if err != nil {

		return nil, err
	}

	userPrefFind, err = upc.IUserPreferencesRepository.FindUserPreferences(context.Background(), id)
	return userPrefFind, nil
}

func (upc *UserPrefController) DeleteUserPreferences() error {
	id, err := helper.GetQueryID(upc.Request)
	if err != nil {
		return err
	}

	if _, err := upc.IUserPreferencesRepository.DeleteUserPreferences(context.Background(), id); err != nil {
		responses.ERROR(upc.Writer, http.StatusInternalServerError, err)
		return err
	}
	responses.JSON(upc.Writer, http.StatusNoContent, "")

	return err
}
