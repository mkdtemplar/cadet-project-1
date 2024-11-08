package controllers

import (
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/repository"
	"cadet-project/pkg/repository/validation"
	"net/http"
)

type Controller struct {
	repository.PG
	Writer  http.ResponseWriter
	Request *http.Request
}

type UserController struct {
	Controller
	interfaces.IUserRepository
}

type VehicleController struct {
	Controller
	interfaces.IUserVehicleRepository
}

type UserPrefController struct {
	Controller
	interfaces.IUserPreferencesRepository
}

type ShipController struct {
	Controller
	interfaces.IShipPortsRepository
	interfaces.IUserPreferencesRepository
	interfaces.IUserRepository
}

type LoginController struct {
	Controller
	interfaces.IShipPortsRepository
	interfaces.IUserRepository
}

type RouteController struct {
	Controller
	interfaces.IUserVehicleRepository
	interfaces.IShipPortsRepository
}

var V validation.Validation
