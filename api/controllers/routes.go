package controllers

import (
	"cadet-project/configurations"
	"cadet-project/middlewares"
	"cadet-project/saml_handler"
	_ "embed"

	"net/http"
)

func (s *Server) InitializeRoutes() {
	configurations.InitConfig("configurations")
	samlSp := saml_handler.AuthorizationRequest()
	app := http.HandlerFunc(s.Home)

	s.Router.Handle("/hello", samlSp.RequireAccount(app))
	s.Router.Handle("/saml/acs", samlSp)
	s.Router.HandleFunc("/userdelete", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteUser))).Methods("DELETE")

	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui")))
	s.Router.PathPrefix("/swaggerui/").Handler(sh)

	// User preferences endpoints
	s.Router.HandleFunc("/userpref", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateUserPreferences))).Methods("POST")
	s.Router.HandleFunc("/userpref", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUserPreferences))).Methods("GET")
	s.Router.HandleFunc("/userpref/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetSingleUserPreference))).Methods("GET")
	s.Router.HandleFunc("/userpref", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUserPreferencesPorts))).Methods("PUT")
	s.Router.HandleFunc("/userpref", middlewares.SetMiddlewareAuthentication(s.DeleteUserPref)).Methods("DELETE")
	s.Router.HandleFunc("/userprefports", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUserPreferencesPorts))).Methods("GET")
}
