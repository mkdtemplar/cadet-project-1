package handlers

import (
	"cadet-project/repository"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	Router *http.ServeMux
	repository.PG
}

func (s *Server) InitializeDB(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	s.PG.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		fmt.Printf("Cannot connect to %s database %s", Dbdriver, DbName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database: %s\n", Dbdriver, DbName)
	}

	s.Router = http.NewServeMux()
	s.InitializeRoutes()

}

func (s *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")

	log.Fatal(http.ListenAndServe(addr, s.Router))
}