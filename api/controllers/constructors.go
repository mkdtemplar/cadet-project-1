package controllers

import (
	"cadet-project/interfaces"
	"cadet-project/repository"
)

func (s *Server) UserRepoFunc() interfaces.IUserRepository {
	return repository.NewUserRepo(s.DB)
}

func (s *Server) UserPrefRepo() interfaces.IUserPreferencesRepository {
	return repository.NewUserPrefRepo(s.DB)
}

func (s *Server) ControllerConstructor() (interfaces.IUserController, interfaces.IUserPrefController) {
	userRepo := s.UserRepoFunc()
	userPrefRepo := s.UserPrefRepo()
	return NewUserController(userRepo), NewUserPrefController(userPrefRepo)
}
