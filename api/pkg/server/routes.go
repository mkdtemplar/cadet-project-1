package server

import (
	"cadet-project/pkg/middlewares"
	"cadet-project/pkg/saml_handler"
	"net/http"
)

func (s *Server) InitializeRoutes() {
	loginController := C.LoginController()
	userController := C.UserController()
	userPrefController := C.UserPrefController()
	shipPortsController := C.ShipPortsController()
	samlSp := saml_handler.AuthorizationRequest()
	app := http.HandlerFunc(loginController.ServeHTTPLogin)

	s.Router.Handle("/login", samlSp.RequireAccount(app))
	s.Router.Handle("/saml/acs", samlSp)

	s.Router.Handle("/", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(userController.ServeHTTPUser)))
	s.Router.Handle("/user_pref", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(userPrefController.ServeHTTPUserPreferences)))
	s.Router.Handle("/user_ports", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(shipPortsController.ServeHTTPShipPorts)))
	s.Router.Handle("/user_pref_ports", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(shipPortsController.ServeHTTPShipPorts)))
}
