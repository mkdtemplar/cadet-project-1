package middlewares_token_validation

import (
	"cadet-project/pkg/responses"
	"errors"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("cookie not found your not authorized"))
			return
		}
		http.SetCookie(w, cookie)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)

	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := ValidateToken(w, r)

		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized token"))
			return
		}
		next.ServeHTTP(w, r)
	}
}
