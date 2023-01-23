package controllers

import (
	"cadet-project/interfaces"
	"cadet-project/repository"
)

type Controller struct {
	interfaces.IUserRepository
	interfaces.IUserPreferencesRepository
	interfaces.IShipPortsRepository
	interfaces.IUserController
	interfaces.IUserPrefController
	interfaces.IShipController
	repository.PG
}
