package saml_handler

import (
	"cadet-project/responses"
	"cadet-project/validation"
	"errors"
	"net/http"

	"github.com/crewjam/saml/samlsp"
)

func Credentials(w http.ResponseWriter, r *http.Request, email string, name string) (string, string) {

	userEmail := samlsp.AttributeFromContext(r.Context(), email)

	if userEmail == "" {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("cannot retrieve user email"))
		return "", ""
	}

	userName := samlsp.AttributeFromContext(r.Context(), name)

	if userName == "" {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("cannot retrieve user name"))
		return "", ""
	}

	err := validation.ValidateUserData(userEmail, userName)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("invalid user email format and/or name format"))
		return "", ""
	}

	return userEmail, userName
}
