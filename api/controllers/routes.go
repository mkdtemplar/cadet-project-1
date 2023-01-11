package controllers

import (
	"cadet-project/middlewares"
	"cadet-project/saml_handler"
	"net/http"
)

func (s *Server) InitializeRoutes() {
	userController, _ := s.ControllerConstructor()
	samlSp := saml_handler.AuthorizationRequest()
	app := http.HandlerFunc(userController.Home)
	s.Router.Handle("/hello", samlSp.RequireAccount(app))
	s.Router.Handle("/saml/acs", samlSp)
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.ServeEndPoints)))

}
