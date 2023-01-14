package controllers

import (
	"cadet-project/middlewares_token_validation"
	"cadet-project/saml_handler"
	"net/http"
)

func (s *Server) InitializeRoutes() {
	userController, _ := s.ControllersConstructor()
	samlSp := saml_handler.AuthorizationRequest()
	app := http.HandlerFunc(userController.Home)
	s.Router.Handle("/hello", samlSp.RequireAccount(app))
	s.Router.Handle("/saml/acs", samlSp)
	s.Router.HandleFunc("/", middlewares_token_validation.SetMiddlewareJSON(middlewares_token_validation.SetMiddlewareAuthentication(s.ServeEndPoints)))

}
