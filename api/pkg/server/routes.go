package server

import (
	"cadet-project/pkg/middlewares"
	"cadet-project/pkg/saml_handler"
	"net/http"
)

func (s *Server) InitializeRoutes() {
	loginController := LoginController.LoginController()
	userController := UserController.UserController()
	userPrefController := UserPrefController.UserPrefController()
	shipPortsController := ShipController.ShipPortsController()
	vehicleController := Vehicle.VehicleController()
	samlSp := saml_handler.AuthorizationRequest()
	app := http.HandlerFunc(loginController.ServeHTTP)

	s.Router.Handle("/login", samlSp.RequireAccount(app))
	s.Router.Handle("/saml/acs", samlSp)

	s.Router.Handle("/", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(userController.ServeHTTP)))
	s.Router.Handle("/user_pref", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(userPrefController.ServeHTTP)))
	s.Router.Handle("/user_ports", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(shipPortsController.ServeHTTP)))
	s.Router.Handle("/user_pref_ports", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(shipPortsController.ServeHTTP)))
	s.Router.Handle("/port_directions", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(shipPortsController.ServeHTTP)))
	s.Router.Handle("/vehicle", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(vehicleController.ServeHTTP)))
	s.Router.Handle("/user_vehicle", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(vehicleController.ServeHTTP)))
	s.Router.Handle("/all_vehicles", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(vehicleController.ServeHTTP)))
}
