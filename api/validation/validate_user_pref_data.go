package validation

import (
	"cadet-project/models"
	"errors"
	"html"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func ValidateUserPref(country string, userId uuid.UUID) error {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]+$`)
	checkId := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

	if country == "" {
		return errors.New("country cannot be empty")
	}
	if checkLetters.MatchString(country) == false {
		return errors.New("country string wrong format")
	}
	if userId == uuid.Nil || checkId.MatchString(userId.String()) == false {
		return errors.New("user id is required or wrong data format user_id must be uuid")
	}

	return nil
}

func NewUserPrefObject(id uuid.UUID, country string, userId uuid.UUID) models.UserPreferences {
	userPref := models.UserPreferences{}
	country = html.EscapeString(strings.TrimSpace(country))

	userPref = models.UserPreferences{
		ID:          id,
		UserCountry: country,
		UserId:      userId,
		Ports:       nil,
	}
	return userPref
}
