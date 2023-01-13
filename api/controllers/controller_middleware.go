package controllers

import (
	"cadet-project/models"
	"cadet-project/responses"
	"cadet-project/validation"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func readBody(w http.ResponseWriter, r *http.Request) models.User {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return models.User{}
	}

	return user
}

func (s *Server) Name(f func(ctx context.Context, usr *models.User) (*models.User, error)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		f = s.IUserRepository.SaveUserDb
	}
}

func (s *Server) TestCreateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := readBody(w, r)

		err := validation.ValidateUserData(user.Email, user.Name)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("invalid user email format"))
			return
		}
		s.IUserRepository.PrepareUserData(user.Email, user.Name)
		if _, err = s.IUserRepository.SaveUserDb(r.Context(), &user); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusCreated, user)
	}
}
