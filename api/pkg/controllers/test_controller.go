package controllers

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/responses"
	"errors"
	"net/http"
)

func (c *Controller) notFound(w http.ResponseWriter) {
	responses.ERROR(w, http.StatusInternalServerError, errors.New("path not found"))
	return
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	c.Writer = w
	c.Request = r

	config.InitConfig("configurations")

	w.Header().Set("content-type", "application/json")

	currentPath := r.URL.Path

	switch currentPath {
	case config.Config.UserDelete:
		c.DeleteUser()
		return
	case config.Config.UserCreate:
		c.CreateIn()
		return

	case config.Config.UserId:
		c.GetUserById()
		return
	default:
		c.notFound(w)
		return
	}

}
func (c *Controller) CreateIn() {
	c.Create(c.Writer, c.Request)
}

func (c *Controller) GetUserById() {
	c.GetId(c.Writer, c.Request, helper.GetQueryID(c.Writer, c.Request))
}

func (c *Controller) DeleteUser() {
	c.Delete(c.Writer, c.Request, helper.GetQueryID(c.Writer, c.Request))
}
