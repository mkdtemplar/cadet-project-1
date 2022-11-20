package server

import (
	"fmt"
	"github.com/crewjam/saml/samlsp"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	s := samlsp.AttributeFromContext(r.Context(), "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name")

	if s == "" {
		return
	}

	_, err := fmt.Fprintf(w, s)
	if err != nil {
		fmt.Fprintf(w, string(rune(http.StatusBadRequest)))
		return
	}
}
