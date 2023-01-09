package validation

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
)

func ValidateUserData(action string, email string, name string) error {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]*$`)
	name = strings.ReplaceAll(name, "\"", "")
	email = strings.ReplaceAll(email, "\"", "")

	switch strings.ToLower(action) {

	case "create":
		if email == "" {
			return errors.New("e-mail is required")
		}
		if _, err := mail.ParseAddress(email); err != nil {
			return errors.New("invalid E-mail format")
		}
		if !checkLetters.MatchString(name) {
			return errors.New("invalid name")
		}
	default:
		if email == "" {
			return errors.New("e-mail is required")
		}
		if _, err := mail.ParseAddress(email); err != nil {
			return errors.New("invalid E-mail format")
		}
		if !checkLetters.MatchString(name) {
			return errors.New("invalid name")
		}
	}
	return nil
}
