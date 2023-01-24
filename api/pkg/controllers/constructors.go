package controllers

import (
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/repository"
)

func (c *Controller) UserRepoConstructor() interfaces.IUserRepository {
	return repository.NewUserRepo(c.DB)
}

func (c *Controller) UserPrefRepoConstructor() interfaces.IUserPreferencesRepository {
	return repository.NewUserPrefRepo(c.DB)
}

func (c *Controller) ShipPortsRepoConstructor() interfaces.IShipPortsRepository {
	return repository.NewShipPortsRepo(c.DB)
}

func (c *Controller) ControllersConstructor() (interfaces.IUserController, interfaces.IUserPrefController, interfaces.IShipController) {
	userRepo := c.UserRepoConstructor()
	userPrefRepo := c.UserPrefRepoConstructor()
	shipPortsRepo := c.ShipPortsRepoConstructor()
	return NewUserController(userRepo, userPrefRepo, shipPortsRepo), NewUserPrefController(userPrefRepo), NewShipPortsController(userRepo, userPrefRepo, shipPortsRepo)
}

func (c *Controller) UserController() interfaces.IUserController {
	userRepo := c.UserRepoConstructor()
	userPrefRepo := c.UserPrefRepoConstructor()
	shipPortsRepo := c.ShipPortsRepoConstructor()
	return NewUserController(userRepo, userPrefRepo, shipPortsRepo)
}

func (c *Controller) UserPrefController() interfaces.IUserPrefController {
	userPrefRepo := c.UserPrefRepoConstructor()
	return NewUserPrefController(userPrefRepo)
}

func (c *Controller) ShipController() interfaces.IShipController {
	userRepo := c.UserRepoConstructor()
	userPrefRepo := c.UserPrefRepoConstructor()
	shipPortsRepo := c.ShipPortsRepoConstructor()
	return NewShipPortsController(userRepo, userPrefRepo, shipPortsRepo)
}

func (c *Controller) Controllers() *Controller {
	userRepo := c.UserRepoConstructor()
	userPrefRepo := c.UserPrefRepoConstructor()
	shipPortsRepo := c.ShipPortsRepoConstructor()
	userController := c.UserController()
	userPrefController := c.UserPrefController()
	shipController := c.ShipController()
	return NewController(userRepo, userPrefRepo, shipPortsRepo, userController, userPrefController, shipController)
}
