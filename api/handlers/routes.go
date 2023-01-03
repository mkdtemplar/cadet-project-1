package handlers

import (
	"cadet-project/configurations"
	"cadet-project/middlewares"
	"cadet-project/repository"
	"cadet-project/saml_handler"
	"net/http"
)

func (s *Server) InitializeRoutes() {
	configurations.InitConfig("configurations")
	samlSp := saml_handler.AuthorizationRequest()
	usrRepo := repository.NewUserRepo(s.DB)
	usrHandlers := NewUserHandler(usrRepo)
	app := http.HandlerFunc(usrHandlers.Home)

	s.Router.Handle("/hello", samlSp.RequireAccount(app))
	s.Router.Handle("/saml/acs", samlSp)

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.ServeEndPoints)))

}
