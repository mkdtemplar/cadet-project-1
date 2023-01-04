package handlers

import (
	"cadet-project/interfaces"
	"cadet-project/repository"
)

func (s *Server) UserRepoFunc() interfaces.IUserRepository {
	var repo = repository.NewUserRepo(s.DB)
	return repo
}

func (s *Server) UserHandlerFunc() interfaces.IUserHandlers {
	userRepo := s.UserRepoFunc()
	return NewUserHandler(userRepo)
}

func (s *Server) UserPrefRepoFunc() interfaces.IUserPreferencesRepository {
	return repository.NewUserPrefRepo(s.DB)
}

func (s *Server) UserPrefHandlerFunc() interfaces.IUserPrefHandlers {
	userPrefRepo := s.UserPrefRepoFunc()
	return NewUserPrefHandler(userPrefRepo)
}
