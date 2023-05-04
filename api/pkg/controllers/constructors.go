package controllers

import (
	"cadet-project/pkg/repository"
)

func (l *LoginController) LoginController() *LoginController {
	userRepo := repository.NewUserRepo()
	shipPortsRepo := repository.NewShipPortsRepo()
	return NewLoginController(userRepo, shipPortsRepo)
}

func (uc *UserController) UserController() *UserController {
	userRepo := repository.NewUserRepo()

	return NewUserController(userRepo)
}

func (upc *UserPrefController) UserPrefController() *UserPrefController {
	userPrefRepo := repository.NewUserPrefRepo()
	return NewUserPrefController(userPrefRepo)
}

func (sp *ShipController) ShipPortsController() *ShipController {
	userRepo := repository.NewUserRepo()
	shipPortsRepo := repository.NewShipPortsRepo()
	userPrefRepo := repository.NewUserPrefRepo()
	return NewShipPortsController(userRepo, userPrefRepo, shipPortsRepo)
}

func (v *VehicleController) VehicleController() *VehicleController {
	vehicleRepo := repository.NewVehicleRepo()

	return NewVehicleController(vehicleRepo)
}

func (r *RouteController) RouteController() *RouteController {
	vehicleRepo := repository.NewVehicleRepo()
	shipPortsRepo := repository.NewShipPortsRepo()
	return NewRouteController(vehicleRepo, shipPortsRepo)
}
