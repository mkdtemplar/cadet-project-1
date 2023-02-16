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

func NewVehicleController(IUserVehicleRepository interfaces.IUserVehicleRepository) *Controller {
	return &Controller{IUserVehicleRepository: IUserVehicleRepository}
}

func (c *Controller) ServeHTTPUserVehicle(w http.ResponseWriter, r *http.Request) {
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
	case config.Config.Vehicle:
		c.ServeVehicleEndPoints(w, r)
		return
	}
}

func (c *Controller) CreateVehicle() (*models.Vehicle, error) {

	vehicle, err := helper.ParseVehicleRequestBody(c.Request)
	if err != nil {
		return nil, errors.New("can not parse request body")
	}

	validateVehicleData := V.ValidateVehicleName(vehicle.Name).ValidateVehicleModel(vehicle.Model).
		ValidateVehicleMileage(vehicle.Mileage).ValidateUserId(vehicle.UserId)

	if validateVehicleData.Err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, validateVehicleData.Err)
		return nil, validateVehicleData.Err
	}

	storeVehicle := repository.NewVehicleObject(generate_id.GenerateID(), vehicle.Name, vehicle.Model, vehicle.Mileage, vehicle.UserId)

	createdVehicle, err := c.IUserVehicleRepository.CreateUserVehicle(context.Background(), &storeVehicle)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return nil, err
	}

	return createdVehicle, nil
}

func (c *Controller) GetVehicleById() (*models.Vehicle, error) {
	id, err := helper.GetQueryID(c.Request)
	if err != nil {
		return nil, err
	}
	return c.IUserVehicleRepository.GetUserVehicleById(context.Background(), id)
}

func (c *Controller) UpdateVehicle() (*models.Vehicle, error) {
	id, err := helper.GetQueryID(c.Request)
	if err != nil {
		return nil, err
	}

	findVehicle, err := c.IUserVehicleRepository.GetUserVehicleById(context.Background(), id)

	if err != nil {
		return &models.Vehicle{}, errors.New("vehicle not found")
	}

	vehicleUpdate, err := helper.ParseVehicleRequestBody(c.Request)
	if err != nil {
		return &models.Vehicle{}, errors.New("cannot parse request body")
	}

	validateVehicle := V.ValidateVehicleName(vehicleUpdate.Name).ValidateVehicleModel(vehicleUpdate.Model).
		ValidateVehicleMileage(vehicleUpdate.Mileage)

	if validateVehicle.Err != nil {
		return &models.Vehicle{}, validateVehicle.Err
	}

	findVehicle, err = c.IUserVehicleRepository.UpdateUserVehicle(context.Background(), vehicleUpdate.Name, vehicleUpdate.Model, vehicleUpdate.Mileage, id)

	findVehicle, err = c.IUserVehicleRepository.GetUserVehicleById(context.Background(), id)

	return findVehicle, nil
}

func (c *Controller) DeleteVehicle() error {
	id, err := helper.GetQueryID(c.Request)
	if err != nil {
		return err
	}
	_, err = c.IUserVehicleRepository.GetUserVehicleById(context.Background(), id)
	if err != nil {
		return errors.New("vehicle do not exist in database")
	}

	if _, err = c.IUserVehicleRepository.DeleteUserVehicle(context.Background(), id); err != nil {
		return errors.New("cannot delete vehicle")
	}

	return nil
}
