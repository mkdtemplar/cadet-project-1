package validation

import (
	"errors"
	"html"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func ValidateUserPref(action string, country string, userId uuid.UUID) error {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]+$`)
	checkId := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

	switch strings.ToLower(action) {
	case "create":
		if country == "" {
			return errors.New("country cannot be empty")
		}
		if checkLetters.MatchString(country) == false {
			return errors.New("country string wrong format")
		}
		if userId == uuid.Nil || checkId.MatchString(userId.String()) == false {
			return errors.New("user id is required or wrong data format user_id must be integer")
		}
	case "update":
		if country == "" {
			return errors.New("country cannot be empty")
		}
		if checkLetters.MatchString(country) == false {
			return errors.New("country string wrong format")
		}
	default:
		if country == "" {
			return errors.New("country cannot be empty")
		}
		if checkLetters.MatchString(country) == false {
			return errors.New("country string wrong format")
		}
		if userId == uuid.Nil || checkId.MatchString(userId.String()) == false {
			return errors.New("user id is required or wrong data format user_id must be integer")
		}
	}
	return nil
}

func ConstructUserPrefObject(country string) {
	country = html.EscapeString(strings.TrimSpace(country))

}
