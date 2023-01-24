package server

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

var C controllers.Controller

func (s *Server) InitializeAPI() {
	config.InitDbConfig("pkg/config")
	C.PG.InitDb()

	s.Router = http.NewServeMux()
	s.InitializeRoutes()

}

func (s *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")

	log.Fatal(http.ListenAndServe(addr, s.Router))
}
