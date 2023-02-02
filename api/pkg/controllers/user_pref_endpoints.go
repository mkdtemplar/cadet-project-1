package controllers

import "C"
import (
	"cadet-project/pkg/responses"
	"fmt"
	"net/http"
)

func (c *Controller) ServeUserPrefEndPoints(w http.ResponseWriter, r *http.Request) {
	c.Writer = w
	c.Request = r
	/*
		id, err := helper.GetQueryID(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}

	*/

	var val interface{}
	var err error
	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 401)
		} else {
			responses.JSON(w, http.StatusOK, val)
		}
	}()

	switch {
	case r.Method == http.MethodPost:
		val, err = c.CreateUserPref()
		return
	case r.Method == http.MethodGet:
		val, err = c.GetUserPrefById()
		return
	case r.Method == http.MethodPatch:
		val, err = c.UpdateUserPref()
		return
	case r.Method == http.MethodDelete:
		err = c.DeleteUserPreferences()
		return
	default:
		err = fmt.Errorf("not found")
		return
	}
}
