package controllers

import (
	"cadet-project/middlewares_token_validation"
	"cadet-project/saml_handler"
	"net/http"
)

func (s *Server) InitializeRoutes() {
	loginController := C.LoginController()
	samlSp := saml_handler.AuthorizationRequest()
	app := http.HandlerFunc(loginController.Login)
	s.Router.Handle("/login", samlSp.RequireAccount(app))
	s.Router.Handle("/saml/acs", samlSp)
	s.Router.HandleFunc("/", middlewares_token_validation.SetMiddlewareJSON(middlewares_token_validation.SetMiddlewareAuthentication(s.ServeEndPoints)))

}
