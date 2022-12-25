package controllers

import (
	"cadet-project/models"
	"cadet-project/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func (s *Server) CreateUserInDb(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}

		user := models.User{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		err = models.ValidateUserData("create", user.Email, user.Name)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("invalid user email format"))
			return
		}

		user.PrepareUserData(user.Email, user.Name)
		userCreated, err := user.SaveUserDb(s.DB)

		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(w, http.StatusCreated, userCreated)
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
		return
	}

}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		paramsID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		user := models.User{}

		_, err = user.DeleteUserDb(s.DB, paramsID)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusNoContent, "")
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
	}

}
