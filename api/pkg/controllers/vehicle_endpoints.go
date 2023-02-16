package controllers

import (
	"cadet-project/pkg/responses"
	"fmt"
	"net/http"
)

func (c *Controller) ServeVehicleEndPoints(w http.ResponseWriter, r *http.Request) {
	c.Writer = w
	c.Request = r

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
		val, err = c.CreateVehicle()
		return
	case r.Method == http.MethodGet:
		val, err = c.GetVehicleById()
		return
	case r.Method == http.MethodPatch:
		val, err = c.UpdateVehicle()
		return
	case r.Method == http.MethodDelete:
		err = c.DeleteVehicle()
		return
	default:
		err = fmt.Errorf("not found")
		return
	}
}
