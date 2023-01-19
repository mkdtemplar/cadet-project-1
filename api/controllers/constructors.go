package controllers

import (
	"cadet-project/interfaces"
	"cadet-project/repository"
)

func (s *Server) UserRepoConstructor() interfaces.IUserRepository {
	return repository.NewUserRepo(s.DB)
}

func (s *Server) UserPrefRepoConstructor() interfaces.IUserPreferencesRepository {
	return repository.NewUserPrefRepo(s.DB)
}

func (s *Server) ShipPortsConstructor() interfaces.IShipPortsRepository {
	return repository.NewShipPortsRepo(s.DB)
}

func (s *Server) ControllersConstructor() (interfaces.IUserController, interfaces.IUserPrefController, interfaces.IShipController) {
	userRepo := s.UserRepoConstructor()
	userPrefRepo := s.UserPrefRepoConstructor()
	shipPortsRepo := s.ShipPortsConstructor()
	return NewUserController(userRepo), NewUserPrefController(userPrefRepo), NewShipPortsController(userPrefRepo, shipPortsRepo)
}
