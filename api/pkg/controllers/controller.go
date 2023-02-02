package controllers

import (
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/repository"
	"cadet-project/pkg/repository/validation"
	"net/http"
)

type Controller struct {
	interfaces.IUserRepository
	interfaces.IUserPreferencesRepository
	interfaces.IShipPortsRepository
	repository.PG
	Writer  http.ResponseWriter
	Request *http.Request
}

var V validation.Validation
