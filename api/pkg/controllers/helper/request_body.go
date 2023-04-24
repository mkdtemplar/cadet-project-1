package helper

import (
	"cadet-project/pkg/models"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func ParseVehicleRequestBody(r *http.Request) (*models.Vehicle, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	vehicle := &models.Vehicle{}
	err = json.Unmarshal(body, &vehicle)

	return vehicle, nil
}

func GetID(r *http.Request) (uuid.UUID, error) {
	queryString := r.URL.Query().Get("id")
	paramsID, err := uuid.Parse(queryString)
	if err != nil {
		return uuid.Nil, err
	}
	return paramsID, nil
}
func GetStartLocation(r *http.Request) string {
	queryString := r.URL.Query().Get("start")
	queryString = cases.Title(language.Und).String(queryString)
	return queryString
}

func GetEndLocation(r *http.Request) string {
	queryString := r.URL.Query().Get("end")
	queryString = cases.Title(language.Und).String(queryString)
	return queryString
}
