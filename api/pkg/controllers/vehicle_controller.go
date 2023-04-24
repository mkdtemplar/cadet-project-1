package controllers

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/repository"
	"cadet-project/pkg/repository/generate_id"
	"cadet-project/pkg/responses"
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

	w.Header().Set("Content-Type", "application/json")

	currentPath := r.URL.Path

	var err error
	var val interface{}

	defer func() {
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		} else {
			responses.JSON(w, http.StatusOK, val)
		}
	}()

	switch currentPath {
	case config.Config.Vehicle:
		v.ServeVehicleEndPoints(w, r)
		return
	case config.Config.UserVehicle:
		val, err = v.GetUserVehicle()
		return

	case config.Config.AllVehicles:
		val, err = v.GetAllVehiclesForUser()
		return

	case config.Config.VehiclesUserId:
		val, err = v.GetVehicleWithUserId()
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
	id, err := helper.GetID(v.Request)
	if err != nil {
		return nil, err
	}
	return v.IUserVehicleRepository.GetUserVehicleById(v.Request.Context(), id)
}

func (v *VehicleController) GetUserVehicle() (*models.User, error) {
	userId, err := helper.GetID(v.Request)

	userVehicle, err := v.IUserVehicleRepository.FindUserVehicle(v.Request.Context(), userId)
	if err != nil {
		return nil, err
	}

	return userVehicle, nil
}

func (v *VehicleController) GetAllVehiclesForUser() ([]*models.Vehicle, error) {
	userId, err := helper.GetID(v.Request)
	if err != nil {
		return nil, err
	}
	vehicles, err := v.IUserVehicleRepository.FindVehiclesForUser(v.Request.Context(), userId)

	return vehicles, nil
}

func (v *VehicleController) GetVehicleWithUserId() ([]*models.Vehicle, error) {

	vehicles, err := v.IUserVehicleRepository.FindVehiclesForUser(v.Request.Context(), UserID)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (v *VehicleController) UpdateVehicle() (*models.Vehicle, error) {
	id, err := helper.GetID(v.Request)
	if err != nil {
		return nil, err
	}

	findVehicle, err := v.IUserVehicleRepository.GetUserVehicleById(v.Request.Context(), id)

	if err != nil && findVehicle == nil {
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
	if err != nil {
		return nil, err
	}

	findVehicle, err = v.IUserVehicleRepository.GetUserVehicleById(v.Request.Context(), id)
	if err != nil {
		return nil, err
	}

	return findVehicle, nil
}

func (v *VehicleController) DeleteVehicle() error {
	id, err := helper.GetID(v.Request)
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
