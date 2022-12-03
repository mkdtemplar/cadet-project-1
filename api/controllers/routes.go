package controllers

import (
	"cadet-project/configurations"
	"cadet-project/saml_handler"
	"net/http"
)

func (s *Server) InitializeRoutes() {
	configurations.InitConfig("configurations")
	samlSp := saml_handler.AuthorizationRequest()
	app := http.HandlerFunc(s.Home)

	s.Router.Handle("/hello", samlSp.RequireAccount(app))
	s.Router.Handle("/saml/acs", samlSp)
	//s.Router.Handle("/hello", middlewares.SetMiddlewareJSON(s.CreateUserInDb)).Methods("POST")

	//s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUserInDb)).Methods("POST")
}
