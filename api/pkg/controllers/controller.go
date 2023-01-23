package controllers

import (
	interfaces2 "cadet-project/pkg/interfaces"
	"cadet-project/pkg/repository"
)

type Controller struct {
	interfaces2.IUserRepository
	interfaces2.IUserPreferencesRepository
	interfaces2.IShipPortsRepository
	interfaces2.IUserController
	interfaces2.IUserPrefController
	interfaces2.IShipController
	repository.PG
}
