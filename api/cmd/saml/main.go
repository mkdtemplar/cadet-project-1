package main

import (
	"github.com/IvanMarkovskiSF/cadet-project/api/saml_handler"
	"github.com/IvanMarkovskiSF/cadet-project/api/server"
	"log"
	"net/http"
)

func main() {
	samlSp := saml_handler.SamlRequest()
	app := http.HandlerFunc(server.Index)
	http.Handle("/hello", samlSp.RequireAccount(app))
	http.Handle("/saml/acs", samlSp)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
