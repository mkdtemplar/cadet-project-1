package controllers

import (
	"cadet-project/formaterrors"
	"cadet-project/models"
	"cadet-project/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func (s *Server) CreateUserPreferences(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

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

	userPrefCreated, err := userPref.SaveUserPreferences(s.DB)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, userPref.ID))
	responses.JSON(w, http.StatusCreated, userPrefCreated)

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

func (s *Server) GetSingleUserPreference(w http.ResponseWriter, r *http.Request) {
	parms := mux.Vars(r)

	paramsID, err := strconv.ParseUint(parms["id"], 10, 32)
	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	userPref := models.UserPreferences{}
	userPreferences, err := userPref.FindOneUserPref(s.DB, uint32(paramsID))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userPreferences)
}

func (s *Server) GetUserPreferencesPorts(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("country")

	userPref := models.UserPreferences{}
	userPreferences, err := userPref.FindUserPrefPorts(s.DB, params)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, userPreferences)
}

func (s *Server) UpdateUserPreferences(w http.ResponseWriter, r *http.Request) {

	paramsID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	userPref := models.UserPreferences{}
	err = s.DB.Debug().Model(models.UserPreferences{}).Where("id = ?", paramsID).Take(&userPref).Error

	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("User preferences not found"))
	}

	body, err := io.ReadAll(r.Body)

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

	userPrefUpdated, err := userPrefUpdate.UpdateUserPref(s.DB)

	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, userPrefUpdated)

}

func (s *Server) DeleteUserPref(w http.ResponseWriter, r *http.Request) {

	paramsID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	userPref := models.UserPreferences{}

	err = s.DB.Debug().Model(models.UserPreferences{}).Where("id = ?", paramsID).Take(&userPref).Error
	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusNotFound, formattedError)
		return
	}

	_, err = userPref.DeleteUserPref(s.DB, uint32(paramsID))

	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusBadRequest, formattedError)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", paramsID))
	responses.JSON(w, http.StatusNoContent, "")
}
