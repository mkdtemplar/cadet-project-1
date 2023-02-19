package controllers

import "C"
import (
	"cadet-project/pkg/responses"
	"fmt"
	"net/http"
)

func (upc *UserPrefController) ServeUserPrefEndPoints(w http.ResponseWriter, r *http.Request) {
	upc.Writer = w
	upc.Request = r

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
		val, err = upc.CreateUserPref()
		return
	case r.Method == http.MethodGet:
		val, err = upc.GetUserPrefById()
		return
	case r.Method == http.MethodPatch:
		val, err = upc.UpdateUserPref()
		return
	case r.Method == http.MethodDelete:
		err = upc.DeleteUserPreferences()
		return
	default:
		err = fmt.Errorf("not found")
		return
	}
}
