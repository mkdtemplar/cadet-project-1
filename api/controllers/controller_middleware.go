package controllers

import (
	"cadet-project/controllers/helper"
	"cadet-project/repository"
	"cadet-project/responses"
	"net/http"
)

func (s *Server) TestCreateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := helper.ParseUserRequestBody(w, r)
		v := repository.Validation{}
		err := v.ValidateUserEmail(user.Email).ValidateUserName(user.Name)
		if err.Err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err.Err)
			return
		}
		s.IUserRepository.PrepareUserData(user.Email, user.Name)
		if _, errRepo := s.IUserRepository.Create(r.Context(), user); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, errRepo)
			return
		}
		responses.JSON(w, http.StatusCreated, user)
	}
}
