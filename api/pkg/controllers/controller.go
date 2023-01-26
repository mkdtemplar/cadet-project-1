package controllers

import (
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/repository"
	"net/http"
)

type Controller struct {
	interfaces.IUserRepository
	interfaces.IUserPreferencesRepository
	interfaces.IShipPortsRepository
	interfaces.IUserController
	interfaces.IUserPrefController
	interfaces.IShipController
	repository.PG
	Writer  http.ResponseWriter
	Request *http.Request
}

func NewControllerForTest(IUserRepository interfaces.IUserRepository, IShipsRepo interfaces.IShipPortsRepository, writer http.ResponseWriter, request *http.Request) *Controller {
	return &Controller{IUserRepository: IUserRepository, IShipPortsRepository: IShipsRepo, Writer: writer, Request: request}
}

func NewController(IUserRepository interfaces.IUserRepository, IUserPreferencesRepository interfaces.IUserPreferencesRepository,
	IShipPortsRepository interfaces.IShipPortsRepository, IUserController interfaces.IUserController,
	IUserPrefController interfaces.IUserPrefController, IShipController interfaces.IShipController) *Controller {
	return &Controller{
		IUserRepository:            IUserRepository,
		IUserPreferencesRepository: IUserPreferencesRepository,
		IShipPortsRepository:       IShipPortsRepository,
		IUserController:            IUserController,
		IUserPrefController:        IUserPrefController,
		IShipController:            IShipController,
	}
}
