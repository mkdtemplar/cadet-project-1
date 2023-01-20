package controllers

import (
	"cadet-project/configurations"
	"cadet-project/interfaces"
	"cadet-project/repository"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
	repository.PG
	interfaces.IUserRepository
	interfaces.IUserPreferencesRepository
	interfaces.IShipPortsRepository
	interfaces.IUserController
	interfaces.IUserPrefController
	interfaces.IShipController
}

func (s *Server) InitializeAPI() {
	configurations.InitDbConfig("configurations")
	s.PG.InitDb()

	s.Router = http.NewServeMux()
	s.InitializeRoutes()

}

func (s *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")

	log.Fatal(http.ListenAndServe(addr, s.Router))
}
