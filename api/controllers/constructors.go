package controllers

import (
	"cadet-project/interfaces"
	"cadet-project/repository"
)

func (c *Controller) UserRepoConstructor() interfaces.IUserRepository {
	return repository.NewUserRepo(c.DB)
}

func (c *Controller) UserPrefRepoConstructor() interfaces.IUserPreferencesRepository {
	return repository.NewUserPrefRepo(c.DB)
}

func (c *Controller) ShipPortsConstructor() interfaces.IShipPortsRepository {
	return repository.NewShipPortsRepo(c.DB)
}

func (c *Controller) ControllersConstructor() (interfaces.IUserController, interfaces.IUserPrefController, interfaces.IShipController) {
	userRepo := c.UserRepoConstructor()
	userPrefRepo := c.UserPrefRepoConstructor()
	shipPortsRepo := c.ShipPortsConstructor()
	return NewUserController(userRepo, userPrefRepo, shipPortsRepo), NewUserPrefController(userPrefRepo), NewShipPortsController(userRepo, userPrefRepo, shipPortsRepo)
}

func (c *Controller) LoginController() interfaces.ILoginController {
	userRepo := c.UserRepoConstructor()
	userPrefRepo := c.UserPrefRepoConstructor()
	shipPortsRepo := c.ShipPortsConstructor()
	return NewLoginController(userRepo, userPrefRepo, shipPortsRepo)
}
