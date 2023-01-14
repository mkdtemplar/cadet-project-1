package controllers

import (
	"cadet-project/repository"
	"cadet-project/responses"
	"errors"
	"net/http"
)

func (s *Server) TestCreateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := ParseUserRequestBody(w, r)

		err := repository.ValidateUserData(user.Email, user.Name)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("invalid user email format"))
			return
		}
		s.IUserRepository.PrepareUserData(user.Email, user.Name)
		if _, err = s.IUserRepository.Create(r.Context(), user); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusCreated, user)
	}
}
