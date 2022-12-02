package server

import (
	"cadet-project/configurations"
	"fmt"
	"github.com/crewjam/saml/samlsp"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}
	}()

	s := samlsp.AttributeFromContext(r.Context(), configurations.Config.Email)

	if s == "" {
		return
	}
	var tokenName *http.Cookie
	tokenName, err = r.Cookie("token")
	if err != nil {
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenName.Value,
		Path:     "/", // Available for all paths
		MaxAge:   5 * 60,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	_, err = fmt.Fprintf(w, "E-mail: %v\n Token: %v\n", s, tokenName.Value) // TODO

}

func GetAllAttributes(w http.ResponseWriter, r *http.Request) {
	s := samlsp.SessionFromContext(r.Context())
	if s == nil {
		return
	}
	sa, ok := s.(samlsp.SessionWithAttributes)
	if !ok {
		return
	}

	fmt.Fprintf(w, "SAML Response: , %v", sa.GetAttributes())
}
