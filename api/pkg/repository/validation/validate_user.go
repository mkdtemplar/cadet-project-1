package validation

import (
	"errors"
	"html"
	"net/mail"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type Validation struct {
	Err error
}

func (v *Validation) Error() string {
	var err string
	return err
}

func (v *Validation) ValidateUserEmail(email string) *Validation {
	email = strings.ReplaceAll(email, "\"", "")
	email = strings.ToLower(email)
	email = html.EscapeString(strings.TrimSpace(email))

	if _, err := mail.ParseAddress(email); err != nil {
		v.Err = errors.New("invalid E-mail format")
		return v
	}
	return v
}

func (v *Validation) ValidateUserName(name string) *Validation {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]*$`)
	if !checkLetters.MatchString(name) {
		v.Err = errors.New("invalid name")
		return v
	}
	return v
}

func (v *Validation) ValidateUserPrefCountry(country string) *Validation {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]+$`)
	if checkLetters.MatchString(country) == false {
		v.Err = errors.New("country string wrong format")
		return v
	}
	return v
}

func (v *Validation) ValidateUserId(userId uuid.UUID) *Validation {
	checkId := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	if userId == uuid.Nil || checkId.MatchString(userId.String()) == false {
		v.Err = errors.New("user id is required or wrong data format user_id must be uuid")
		return v
	}
	return v
}
