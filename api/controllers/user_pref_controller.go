package controllers

import (
	"cadet-project/formaterrors"
	"cadet-project/models"
	"cadet-project/responses"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (s *Server) CreateUserPreferences(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userPref := models.UserPreferences{}

	err = json.Unmarshal(body, &userPref)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userPref.PrepareUserPref()
	err = userPref.ValidateUserPref()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, userPref.ID))
	responses.JSON(w, http.StatusCreated, userPref)

}

func (s *Server) GetUserPreferences(w http.ResponseWriter, r *http.Request) {
	userPref := models.UserPreferences{}

	userPreferences, err := userPref.FindAllUserPref(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, userPreferences)
}

func (s *Server) UpdateUserPreferences(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramsID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = models.TokenValid(r)

	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	userPref := models.UserPreferences{}
	err = s.DB.Debug().Model(models.UserPreferences{}).Where("id = ?", paramsID).Take(&userPref).Error

	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("User preferences not found"))
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userPrefUpdate := models.UserPreferences{}

	err = json.Unmarshal(body, &userPrefUpdate)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userPrefUpdate.ID = userPref.ID

	userPrefUpdated, err := userPref.UpdateAPost(s.DB)

	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, userPrefUpdated)

}
