package controllers

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/repository"
	"cadet-project/pkg/repository/generate_id"
	"cadet-project/pkg/responses"
	"context"
	"errors"
	"net/http"
)

func NewVehicleController(IUserVehicleRepository interfaces.IUserVehicleRepository) *VehicleController {
	return &VehicleController{IUserVehicleRepository: IUserVehicleRepository}
}

func (v *VehicleController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	v.Request = r
	v.Writer = w
	config.InitConfig("configurations")

	w.Header().Set("content-type", "application/json")

	currentPath := r.URL.Path

	var err error
	var val interface{}

	defer func() {
		if err != nil {
			responses.JSON(w, http.StatusUnprocessableEntity, err)
		} else {
			responses.JSON(w, http.StatusOK, val)
		}
	}()

	switch currentPath {
	case config.Config.Vehicle:
		v.ServeVehicleEndPoints(w, r)
		return
	}
}

func (v *VehicleController) CreateVehicle() (*models.Vehicle, error) {

	vehicle, err := helper.ParseVehicleRequestBody(v.Request)
	if err != nil {
		return nil, errors.New("can not parse request body")
	}

	validateVehicleData := V.ValidateVehicleName(vehicle.Name).ValidateVehicleModel(vehicle.Model).
		ValidateVehicleMileage(vehicle.Mileage).ValidateUserId(vehicle.UserId)

	if validateVehicleData.Err != nil {
		responses.ERROR(v.Writer, http.StatusUnprocessableEntity, validateVehicleData.Err)
		return nil, validateVehicleData.Err
	}

	storeVehicle := repository.NewVehicleObject(generate_id.GenerateID(), vehicle.Name, vehicle.Model, vehicle.Mileage, vehicle.UserId)

	createdVehicle, err := v.IUserVehicleRepository.CreateUserVehicle(v.Request.Context(), &storeVehicle)
	if err != nil {
		responses.ERROR(v.Writer, http.StatusUnprocessableEntity, err)
		return nil, err
	}

	return createdVehicle, nil
}

func (v *VehicleController) GetVehicleById() (*models.Vehicle, error) {
	id, err := helper.GetQueryID(v.Request)
	if err != nil {
		return nil, err
	}
	return v.IUserVehicleRepository.GetUserVehicleById(v.Request.Context(), id)
}

func (v *VehicleController) UpdateVehicle() (*models.Vehicle, error) {
	id, err := helper.GetQueryID(v.Request)
	if err != nil {
		return nil, err
	}

	findVehicle, err := v.IUserVehicleRepository.GetUserVehicleById(context.Background(), id)

	if err != nil {
		return &models.Vehicle{}, errors.New("vehicle not found")
	}

	vehicleUpdate, err := helper.ParseVehicleRequestBody(v.Request)
	if err != nil {
		return &models.Vehicle{}, errors.New("cannot parse request body")
	}

	validateVehicle := V.ValidateVehicleName(vehicleUpdate.Name).ValidateVehicleModel(vehicleUpdate.Model).
		ValidateVehicleMileage(vehicleUpdate.Mileage)

	if validateVehicle.Err != nil {
		return &models.Vehicle{}, validateVehicle.Err
	}

	findVehicle, err = v.IUserVehicleRepository.UpdateUserVehicle(v.Request.Context(), vehicleUpdate.Name, vehicleUpdate.Model, vehicleUpdate.Mileage, id)

	findVehicle, err = v.IUserVehicleRepository.GetUserVehicleById(v.Request.Context(), id)

	return findVehicle, nil
}

func (v *VehicleController) DeleteVehicle() error {
	id, err := helper.GetQueryID(v.Request)
	if err != nil {
		return err
	}
	_, err = v.IUserVehicleRepository.GetUserVehicleById(v.Request.Context(), id)
	if err != nil {
		return errors.New("vehicle do not exist in database")
	}

	if _, err = v.IUserVehicleRepository.DeleteUserVehicle(v.Request.Context(), id); err != nil {
		return errors.New("cannot delete vehicle")
	}

	return nil
}
