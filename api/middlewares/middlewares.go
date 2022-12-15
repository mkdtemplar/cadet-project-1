package middlewares

import (
	"cadet-project/models"
	"cadet-project/responses"
	"errors"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &models.Cookie)
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := models.TokenValid(w, r)

		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized token"))
			return
		}
		next(w, r)
	}

}
