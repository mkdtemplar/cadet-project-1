package handlers

import (
	"cadet-project/middlewares"
	"cadet-project/repository"
	"cadet-project/saml_handler"
	"net/http"
)

func (s *Server) InitializeRoutes() {
	var UserRepo = repository.NewUserRepo(s.DB)
	var UserHandlers = NewUserHandler(UserRepo)
	samlSp := saml_handler.AuthorizationRequest()
	app := http.HandlerFunc(UserHandlers.Home)
	s.Router.Handle("/hello", samlSp.RequireAccount(app))
	s.Router.Handle("/saml/acs", samlSp)
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.ServeEndPoints)))

}
