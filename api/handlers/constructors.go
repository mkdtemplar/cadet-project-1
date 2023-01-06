package handlers

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

func (s *Server) HandlersConstructor() (interfaces.IUserHandlers, interfaces.IUserPrefHandlers) {
	userRepo := s.UserRepoFunc()
	userPrefRepo := s.UserPrefRepo()
	return NewServerUser(userRepo), NewServerUserPref(userPrefRepo)
}
