package helper

import (
	"cadet-project/models"
	"cadet-project/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func ParseUserRequestBody(w http.ResponseWriter, r *http.Request) *models.User {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := &models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return &models.User{}
	}

	return user
}

func ParseUserPrefRequestBody(w http.ResponseWriter, r *http.Request) *models.UserPreferences {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	userPref := &models.UserPreferences{}
	err = json.Unmarshal(body, &userPref)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return &models.UserPreferences{}
	}

	return userPref
}

func GetQueryID(w http.ResponseWriter, r *http.Request) uuid.UUID {
	queryString := r.URL.Query().Get("id")
	paramsID, err := uuid.Parse(queryString)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in id format must be uuid"))
		return uuid.Nil
	}
	return paramsID
}

func GetQueryUserID(w http.ResponseWriter, r *http.Request) uuid.UUID {
	queryString := r.URL.Query().Get("user_id")
	paramsID, err := uuid.Parse(queryString)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in user_id format must be uuid"))
		return uuid.Nil
	}
	return paramsID
}
