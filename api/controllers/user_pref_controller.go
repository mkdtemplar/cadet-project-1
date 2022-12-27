package controllers

import (
	"cadet-project/models"
	"cadet-project/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (s *Server) CreateUserPreferences(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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
		fmt.Println(userPref.UserId)
		err = userPref.ValidateUserPref("create")
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		userPref.ConstructUserPrefObject(userPref.Country, userPref.UserId)

		userPrefCreated, err := userPref.SaveUserPreferences(s.DB)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusCreated, userPrefCreated)
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
	}

}

func (s *Server) GetUserPreference(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		params := r.URL.Query().Get("user_id")

		paramsID, err := strconv.ParseUint(params, 10, 32)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in user_id format must be integer"))
			return
		}

		userPref := models.UserPreferences{}
		userPreferences, err := userPref.FindUserPreferences(s.DB, uint32(paramsID))
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, userPreferences)
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
	}

}

func (s *Server) GetUserPorts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		params := r.URL.Query().Get("user_id")

		paramsID, err := strconv.ParseUint(params, 10, 32)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in user_id format must be integer"))
			return
		}

		userPref := models.UserPreferences{}
		userPreferences, err := userPref.FindUserPreferences(s.DB, uint32(paramsID))
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		userPorts, err := userPreferences.FindUserPrefPorts(s.DB, userPreferences.Country)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(w, http.StatusOK, userPorts)
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
	}
}

func (s *Server) UpdateUserPreferences(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		paramsID, err := strconv.ParseUint(r.URL.Query().Get("user_id"), 10, 32)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in user_id format must be integer"))
			return
		}

		userPref := models.UserPreferences{}
		err = s.DB.Debug().Model(models.UserPreferences{}).Where("user_id = ?", paramsID).Take(&userPref).Error

		if err != nil {
			responses.ERROR(w, http.StatusNotFound, errors.New("user preferences not found"))
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

		err = userPrefUpdate.ValidateUserPref("update")
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		userPrefUpdate.UserId = userPref.UserId

		userPrefUpdated, err := userPrefUpdate.UpdateUserPref(s.DB, uint32(paramsID))

		if err != nil {

			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, userPrefUpdated)
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
	}
}

func (s *Server) DeleteUserPref(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		paramsID, err := strconv.ParseUint(r.URL.Query().Get("user_id"), 10, 32)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in user_id format must be integer"))
			return
		}

		userPref := models.UserPreferences{}

		err = s.DB.Debug().Model(models.UserPreferences{}).Where("user_id = ?", paramsID).Take(&userPref).Error
		if err != nil {

			responses.ERROR(w, http.StatusNotFound, err)
			return
		}

		_, err = userPref.DeleteUserPreferences(s.DB, uint32(paramsID))

		if err != nil {

			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusNoContent, "")
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
	}

}
