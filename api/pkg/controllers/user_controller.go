package controllers

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/responses"
	"context"
	"errors"
	"fmt"
	"net/http"
)

func NewUserController(IUserRepository interfaces.IUserRepository) *UserController {
	return &UserController{IUserRepository: IUserRepository}
}

func (uc *UserController) notFound(w http.ResponseWriter) {
	responses.ERROR(w, http.StatusInternalServerError, errors.New("path not found"))
	return
}

func (uc *UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uc.Writer = w
	uc.Request = r

	w.Header().Set("content-type", "application/json")

	currentPath := r.URL.Path

	var err error
	var val interface{}

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 401)
		} else {
			responses.JSON(w, http.StatusOK, val)
		}
	}()

	switch currentPath {
	case config.Config.UserDelete:
		err = uc.DeleteUser()
		return
	case config.Config.UserCreate:
		val, err = uc.CreateIn()
		return

	case config.Config.UserId:
		val, err = uc.GetUserById()
		return
	default:
		err = fmt.Errorf("not found")
		return
	}

}
func (uc *UserController) CreateIn() (*models.User, error) {
	user, err := helper.ParseUserRequestBody(uc.Request)
	if err != nil {
		return nil, err
	}
	checkCredentials := V.ValidateUserEmail(user.Email).ValidateUserName(user.Name)

	if checkCredentials.Err != nil {
		responses.ERROR(uc.Writer, http.StatusUnprocessableEntity, checkCredentials.Err)
		return nil, err
	}
	_, err = uc.IUserRepository.Create(uc.Request.Context(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserController) GetUserById() (*models.User, error) {
	id, err := helper.GetQueryID(uc.Request)
	if err != nil {
		return nil, err
	}
	return uc.IUserRepository.GetById(context.Background(), id)
}

func (uc *UserController) DeleteUser() error {
	id, err := helper.GetQueryID(uc.Request)
	if err != nil {
		return err
	}

	_, err = uc.IUserRepository.Delete(context.Background(), id)
	return err

}
