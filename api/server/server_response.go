package server

import (
	"cadet-project/configurations"
	"fmt"
	"github.com/crewjam/saml/samlsp"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	config, err := configurations.LoadConfig("./configurations")
	if err != nil {
		log.Fatalln("cannot load configurations")
	}

	s := samlsp.AttributeFromContext(r.Context(), config.Email)

	if s == "" {
		return
	}

	_, err = fmt.Fprintf(w, s)
	if err != nil {
		fmt.Fprintf(w, string(rune(http.StatusBadRequest)))
		return
	}
}
