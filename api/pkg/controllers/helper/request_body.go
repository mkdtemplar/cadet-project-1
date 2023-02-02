package helper

import (
	"cadet-project/pkg/models"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func ParseUserRequestBody(r *http.Request) (*models.User, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	user.Clean()
	return user, nil
}

func ParseUserPrefRequestBody(r *http.Request) (*models.UserPreferences, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	userPref := &models.UserPreferences{}
	err = json.Unmarshal(body, &userPref)
	if err != nil {
		return &models.UserPreferences{}, err
	}

	return userPref, nil
}

func GetQueryID(r *http.Request) (uuid.UUID, error) {
	queryString := r.URL.Query().Get("id")
	paramsID, err := uuid.Parse(queryString)
	if err != nil {
		return uuid.Nil, err
	}
	return paramsID, nil
}

func GetQueryUserID(r *http.Request) (uuid.UUID, error) {
	queryString := r.URL.Query().Get("user_id")
	paramsID, err := uuid.Parse(queryString)
	if err != nil {
		return uuid.Nil, err
	}
	return paramsID, err
}
