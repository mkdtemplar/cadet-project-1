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

func (c *Controller) VehicleRepoConstructor() interfaces.IUserVehicleRepository {
	return repository.NewVehicleRepo(c.DB)
}

func (c *Controller) LoginController() *Controller {
	userRepo := c.UserRepoConstructor()
	shipPortsRepo := c.ShipPortsRepoConstructor()
	return NewLoginController(userRepo, shipPortsRepo)
}

func (c *Controller) UserController() *Controller {
	userRepo := c.UserRepoConstructor()

	return NewUserController(userRepo)
}

func (c *Controller) UserPrefController() *Controller {
	userPrefRepo := c.UserPrefRepoConstructor()
	return NewUserPrefController(userPrefRepo)
}

func (c *Controller) ShipPortsController() *Controller {
	userRepo := c.UserRepoConstructor()
	shipPortsRepo := c.ShipPortsRepoConstructor()
	userPrefRepo := c.UserPrefRepoConstructor()
	return NewShipPortsController(userRepo, userPrefRepo, shipPortsRepo)
}

func (c *Controller) VehicleController() *Controller {
	vehicleRepo := c.VehicleRepoConstructor()
	return NewVehicleController(vehicleRepo)
}
