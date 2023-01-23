package controllers

import (
	interfaces2 "cadet-project/pkg/interfaces"
	repository2 "cadet-project/pkg/repository"
)

func (c *Controller) UserRepoConstructor() interfaces2.IUserRepository {
	return repository2.NewUserRepo(c.DB)
}

func (c *Controller) UserPrefRepoConstructor() interfaces2.IUserPreferencesRepository {
	return repository2.NewUserPrefRepo(c.DB)
}

func (c *Controller) ShipPortsConstructor() interfaces2.IShipPortsRepository {
	return repository2.NewShipPortsRepo(c.DB)
}

func (c *Controller) ControllersConstructor() (interfaces2.IUserController, interfaces2.IUserPrefController, interfaces2.IShipController) {
	userRepo := c.UserRepoConstructor()
	userPrefRepo := c.UserPrefRepoConstructor()
	shipPortsRepo := c.ShipPortsConstructor()
	return NewUserController(userRepo, userPrefRepo, shipPortsRepo), NewUserPrefController(userPrefRepo), NewShipPortsController(userRepo, userPrefRepo, shipPortsRepo)
}

func (c *Controller) LoginController() interfaces2.ILoginController {
	userRepo := c.UserRepoConstructor()
	userPrefRepo := c.UserPrefRepoConstructor()
	shipPortsRepo := c.ShipPortsConstructor()
	return NewLoginController(userRepo, userPrefRepo, shipPortsRepo)
}
