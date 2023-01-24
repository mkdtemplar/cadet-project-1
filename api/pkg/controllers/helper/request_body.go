package helper

import (
	models2 "cadet-project/pkg/models"
	"cadet-project/pkg/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func ParseUserRequestBody(w http.ResponseWriter, r *http.Request) *models2.User {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := &models2.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return &models2.User{}
	}
	user.Clean()
	return user
}

func ParseUserPrefRequestBody(w http.ResponseWriter, r *http.Request) *models2.UserPreferences {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	userPref := &models2.UserPreferences{}
	err = json.Unmarshal(body, &userPref)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return &models2.UserPreferences{}
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
