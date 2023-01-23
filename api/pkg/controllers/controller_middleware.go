package controllers

import (
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/repository/validation"
	"cadet-project/pkg/responses"
	"net/http"
)

func (c *Controller) TestCreateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := helper.ParseUserRequestBody(w, r)
		v := validation.Validation{}
		err := v.ValidateUserEmail(user.Email).ValidateUserName(user.Name)
		if err.Err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err.Err)
			return
		}
		if _, errRepo := c.IUserRepository.Create(r.Context(), user); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, errRepo)
			return
		}
		responses.JSON(w, http.StatusCreated, user)
	}
}
