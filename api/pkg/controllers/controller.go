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

func NewController(IUserRepository interfaces.IUserRepository, IUserPreferencesRepository interfaces.IUserPreferencesRepository, IShipPortsRepository interfaces.IShipPortsRepository, IUserController interfaces.IUserController, IUserPrefController interfaces.IUserPrefController, IShipController interfaces.IShipController) *Controller {
	return &Controller{IUserRepository: IUserRepository, IUserPreferencesRepository: IUserPreferencesRepository, IShipPortsRepository: IShipPortsRepository, IUserController: IUserController, IUserPrefController: IUserPrefController, IShipController: IShipController}
}
