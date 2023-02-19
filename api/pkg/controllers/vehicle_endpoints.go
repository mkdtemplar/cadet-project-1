package controllers

import (
	"cadet-project/pkg/responses"
	"fmt"
	"net/http"
)

func (v *VehicleController) ServeVehicleEndPoints(w http.ResponseWriter, r *http.Request) {
	v.Writer = w
	v.Request = r

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
		val, err = v.CreateVehicle()
		return
	case r.Method == http.MethodGet:
		val, err = v.GetVehicleById()
		return
	case r.Method == http.MethodPatch:
		val, err = v.UpdateVehicle()
		return
	case r.Method == http.MethodDelete:
		err = v.DeleteVehicle()
		return
	default:
		err = fmt.Errorf("not found")
		return
	}
}
