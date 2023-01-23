package saml_handler

import (
	"cadet-project/pkg/repository/validation"
	"cadet-project/pkg/responses"
	"errors"
	"net/http"

	"github.com/crewjam/saml/samlsp"
)

func Credentials(w http.ResponseWriter, r *http.Request, email string, name string) (string, string) {
	v := validation.Validation{}
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

	err := v.ValidateUserEmail(userEmail).ValidateUserName(userName)
	if err.Err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err.Err)
		return "", ""
	}

	return userEmail, userName
}
