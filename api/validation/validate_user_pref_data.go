package validation

import (
	"errors"
	"html"
	"regexp"
	"strconv"
	"strings"
)

func ValidateUserPref(action string, country string, userId uint32) error {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]+$`)
	checkNumber := regexp.MustCompile(`^[0-9]+$`)

	switch strings.ToLower(action) {
	case "create":
		if country == "" {
			return errors.New("country cannot be empty")
		}
		if checkLetters.MatchString(country) == false {
			return errors.New("country string wrong format")
		}
		if userId < 1 || checkNumber.MatchString(strconv.Itoa(int(userId))) == false {
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
		if userId < 1 || checkNumber.MatchString(strconv.Itoa(int(userId))) == false {
			return errors.New("user id is required or wrong data format user_id must be integer")
		}
	}
	return nil
}

func ConstructUserPrefObject(country string) {
	country = html.EscapeString(strings.TrimSpace(country))

}
