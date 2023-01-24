package controllers

import (
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/repository"
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
