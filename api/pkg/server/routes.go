package server

import (
	"cadet-project/pkg/middlewares"
	"cadet-project/pkg/saml_handler"
	"net/http"
)

func (s *Server) InitializeRoutes() {
	loginController := C.LoginController()
	samlSp := saml_handler.AuthorizationRequest()
	app := http.HandlerFunc(loginController.Login)
	s.Router.Handle("/login", samlSp.RequireAccount(app))
	s.Router.Handle("/saml/acs", samlSp)
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.ServeEndPoints)))

}
