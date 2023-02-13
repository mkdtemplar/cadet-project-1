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

func NewUserController(IUserRepository interfaces.IUserRepository) *Controller {
	return &Controller{IUserRepository: IUserRepository}
}

func (c *Controller) notFound(w http.ResponseWriter) {
	responses.ERROR(w, http.StatusInternalServerError, errors.New("path not found"))
	return
}

func (c *Controller) ServeHTTPUser(w http.ResponseWriter, r *http.Request) {
	c.Writer = w
	c.Request = r

	config.InitConfig("configurations")

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
		err = c.DeleteUser()
		return
	case config.Config.UserCreate:
		val, err = c.CreateIn()
		return

	case config.Config.UserId:
		val, err = c.GetUserById()
		return
	default:
		err = fmt.Errorf("not found")
		return
	}

}
func (c *Controller) CreateIn() (*models.User, error) {
	user, err := helper.ParseUserRequestBody(c.Request)
	if err != nil {
		return nil, err
	}
	checkCredentials := V.ValidateUserEmail(user.Email).ValidateUserName(user.Name)

	if checkCredentials.Err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, checkCredentials.Err)
		return nil, err
	}
	_, err = c.IUserRepository.Create(c.Request.Context(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Controller) GetUserById() (*models.User, error) {
	id, err := helper.GetQueryID(c.Request)
	if err != nil {
		return nil, err
	}
	return c.IUserRepository.GetById(context.Background(), id)
}

func (c *Controller) DeleteUser() error {
	id, err := helper.GetQueryID(c.Request)
	if err != nil {
		return err
	}

	_, err = c.IUserRepository.Delete(context.Background(), id)
	return err

}
