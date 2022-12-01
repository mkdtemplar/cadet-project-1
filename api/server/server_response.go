package server

import (
	"cadet-project/configurations"
	"fmt"
	"github.com/crewjam/saml/samlsp"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Before running docker-compose change path to ./configurations // make const
	config, err := configurations.LoadConfig("./api/configurations")
	if err != nil {
		log.Fatalln("cannot load configurations")
	}

	s := samlsp.AttributeFromContext(r.Context(), config.Email)

	if s == "" {
		return
	}

	tokenValue, err := r.Cookie("token")
	if err != nil {
		log.Fatalf("Error occured while reading cookie")
	}
	_, err = fmt.Fprintf(w, "E-mail: %v\n Token: %v\n", s, tokenValue.Value)

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
