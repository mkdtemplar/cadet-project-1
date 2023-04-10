package server

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers"
	"cadet-project/pkg/repository"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

var Vehicle controllers.VehicleController
var UserController controllers.UserController
var UserPrefController controllers.UserPrefController
var ShipController controllers.ShipController
var LoginController controllers.LoginController
var RouteController controllers.RouteController

func (s *Server) InitializeServer() {
	config.InitDbConfig("pkg/config")
	repository.InitDb()

	s.Router = http.NewServeMux()
	s.InitializeRoutes()

}

func (s *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")

	log.Fatal(http.ListenAndServe(addr, s.Router))
}
