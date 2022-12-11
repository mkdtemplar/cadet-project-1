package controllers

import (
	"cadet-project/configurations"
	"cadet-project/middlewares"
	"cadet-project/saml_handler"
	"net/http"
)

func (s *Server) InitializeRoutes() {
	configurations.InitConfig("configurations")
	samlSp := saml_handler.AuthorizationRequest()
	app := http.HandlerFunc(s.Home)

	s.Router.Handle("/hello", samlSp.RequireAccount(app))
	s.Router.Handle("/saml/acs", samlSp)
	s.Router.HandleFunc("/userdelete", middlewares.SetMiddlewareJSON(s.DeleteUser)).Methods("DELETE")

	// User preferences endpoints
	s.Router.HandleFunc("/userpref", middlewares.SetMiddlewareJSON(s.CreateUserPreferences)).Methods("POST")
	s.Router.HandleFunc("/userpref", middlewares.SetMiddlewareJSON(s.GetUserPreferences)).Methods("GET")
	s.Router.HandleFunc("/userpref/{id}", middlewares.SetMiddlewareJSON(s.GetSingleUserPreference)).Methods("GET")
	s.Router.HandleFunc("/userpref", middlewares.SetMiddlewareJSON(s.UpdateUserPreferences)).Methods("PUT")
	s.Router.HandleFunc("/userpref", middlewares.SetMiddlewareJSON(s.DeleteUserPref)).Methods("DELETE")
}
