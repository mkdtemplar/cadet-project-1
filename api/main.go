package main

import (
	"cadet-project/configurations"
	"cadet-project/saml_handler"
	"cadet-project/server"
	"log"
	"net/http"
)

func main() {

	configurations.InitConfig("configurations")
	samlSp := saml_handler.AuthorizationRequest()
	app := http.HandlerFunc(server.Index)
	http.Handle("/hello", samlSp.RequireAccount(app))
	http.Handle("/saml/acs", samlSp)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
