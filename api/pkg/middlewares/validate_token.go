package middlewares

import (
	"cadet-project/pkg/models"
	"cadet-project/pkg/responses"
	"errors"
	"net/http"
)

func ExtractToken(r *http.Request) string {
	tokenName, err := r.Cookie("token")

	if err != nil {
		return ""
	}
	return tokenName.Value
}

func ValidateToken(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("token")
	if err != nil {
		return err
	}
	sessionToken := cookie.Value
	userSession, exists := models.GetSession()[sessionToken]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("token not present in session"))
		return errors.New("invalid token")
	}

	if userSession.IsExpired() {
		delete(models.GetSession(), sessionToken)
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return errors.New("unauthorized")
	}

	return nil
}
