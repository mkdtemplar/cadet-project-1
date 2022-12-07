package controllers

import (
	"cadet-project/formaterrors"
	"cadet-project/models"
	"cadet-project/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (s *Server) CreateUserInDb(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.PrepareUserData()
	err = user.ValidateUserData("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("invalid user email format"))
	}

	userCreated, err := user.SaveUserDb(s.DB)

	if err != nil {

		formattedError := formaterrors.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}
