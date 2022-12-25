package controllers

import (
	"cadet-project/models"
	"cadet-project/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func trimSpaces(param string) string {
	re := regexp.MustCompile(`\s+`)
	out := re.ReplaceAllString(param, " ")
	out = strings.TrimSpace(out)

	return out
}

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

		err = userPref.ValidateUserPref(userPref.Country, userPref.UserId)
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
		params := r.URL.Query().Get("id")

		paramsID, err := strconv.ParseUint(params, 10, 32)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, err)
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
		params := r.URL.Query().Get("country")
		params = trimSpaces(params)
		strings.Replace(params, " ", "%20", -1)

		userPref := models.UserPreferences{}

		userPreferences, err := userPref.FindUserPrefPorts(s.DB, params)

		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(w, http.StatusOK, userPreferences)
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
	}

}

func (s *Server) UpdateUserPreferences(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		paramsID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		userPref := models.UserPreferences{}
		err = s.DB.Debug().Model(models.UserPreferences{}).Where("id = ?", paramsID).Take(&userPref).Error

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
		userPrefUpdate.ID = userPref.ID

		err = userPrefUpdate.ValidateUserPref(userPrefUpdate.Country, userPrefUpdate.UserId)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		userPrefUpdated, err := userPrefUpdate.UpdateUserPref(s.DB)

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
		paramsID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		userPref := models.UserPreferences{}

		err = s.DB.Debug().Model(models.UserPreferences{}).Where("id = ?", paramsID).Take(&userPref).Error
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
