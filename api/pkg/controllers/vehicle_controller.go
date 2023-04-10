package controllers

import (
	"cadet-project/google_API"
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/repository"
	"cadet-project/pkg/repository/generate_id"
	"cadet-project/pkg/responses"
	"errors"
	"net/http"
)

func NewVehicleController(IUserVehicleRepository interfaces.IUserVehicleRepository, IShipPortsRepository interfaces.IShipPortsRepository, IUserRepository interfaces.IUserRepository) *VehicleController {
	return &VehicleController{IUserVehicleRepository: IUserVehicleRepository, IShipPortsRepository: IShipPortsRepository, IUserRepository: IUserRepository}
}

func (v *VehicleController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	v.Request = r
	v.Writer = w
	config.InitConfig("configurations")

	w.Header().Set("Content-Type", "application/json")

	currentPath := r.URL.Path

	var err error
	var val interface{}

	defer func() {
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		} else {
			responses.JSON(w, http.StatusOK, val)
		}
	}()

	switch currentPath {
	case config.Config.Vehicle:
		v.ServeVehicleEndPoints(w, r)
		return
	case config.Config.UserVehicle:
		val, err = v.GetUserVehicle()
		return

	case config.Config.AllVehicles:
		val, err = v.GetAllVehiclesForUser()
		return

	case config.Config.VehiclesUserId:
		val, err = v.GetVehicleWithUserId()
		return

	case config.Config.PortName:
		val, err = v.GetDirections()
		return
	}
}

func (v *VehicleController) CreateVehicle() (*models.Vehicle, error) {

	vehicle, err := helper.ParseVehicleRequestBody(v.Request)
	if err != nil {
		return nil, errors.New("can not parse request body")
	}

	validateVehicleData := V.ValidateVehicleName(vehicle.Name).ValidateVehicleModel(vehicle.Model).
		ValidateVehicleMileage(vehicle.Mileage).ValidateUserId(vehicle.UserId)

	if validateVehicleData.Err != nil {
		responses.ERROR(v.Writer, http.StatusUnprocessableEntity, validateVehicleData.Err)
		return nil, validateVehicleData.Err
	}

	storeVehicle := repository.NewVehicleObject(generate_id.GenerateID(), vehicle.Name, vehicle.Model, vehicle.Mileage, vehicle.UserId)

	createdVehicle, err := v.IUserVehicleRepository.CreateUserVehicle(v.Request.Context(), &storeVehicle)
	if err != nil {
		responses.ERROR(v.Writer, http.StatusUnprocessableEntity, err)
		return nil, err
	}

	return createdVehicle, nil
}

func (v *VehicleController) GetVehicleById() (*models.Vehicle, error) {
	id, err := helper.GetQueryID(v.Request)
	if err != nil {
		return nil, err
	}
	return v.IUserVehicleRepository.GetUserVehicleById(v.Request.Context(), id)
}

func (v *VehicleController) GetUserVehicle() (*models.User, error) {
	userId, err := helper.GetQueryID(v.Request)

	userVehicle, err := v.IUserVehicleRepository.FindUserVehicle(v.Request.Context(), userId)
	if err != nil {
		return nil, err
	}

	return userVehicle, nil
}

func (v *VehicleController) GetAllVehiclesForUser() ([]*models.Vehicle, error) {
	userId, err := helper.GetQueryID(v.Request)
	if err != nil {
		return nil, err
	}
	vehicles, err := v.IUserVehicleRepository.FindVehiclesForUser(v.Request.Context(), userId)

	return vehicles, nil
}

func (v *VehicleController) GetVehicleWithUserId() ([]*models.Vehicle, error) {

	vehicles, err := v.IUserVehicleRepository.FindVehiclesForUser(v.Request.Context(), UserID)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (v *VehicleController) UpdateVehicle() (*models.Vehicle, error) {
	id, err := helper.GetQueryID(v.Request)
	if err != nil {
		return nil, err
	}

	findVehicle, err := v.IUserVehicleRepository.GetUserVehicleById(v.Request.Context(), id)

	if err != nil && findVehicle == nil {
		return &models.Vehicle{}, errors.New("vehicle not found")
	}

	vehicleUpdate, err := helper.ParseVehicleRequestBody(v.Request)
	if err != nil {
		return &models.Vehicle{}, errors.New("cannot parse request body")
	}

	validateVehicle := V.ValidateVehicleName(vehicleUpdate.Name).ValidateVehicleModel(vehicleUpdate.Model).
		ValidateVehicleMileage(vehicleUpdate.Mileage)

	if validateVehicle.Err != nil {
		return &models.Vehicle{}, validateVehicle.Err
	}

	findVehicle, err = v.IUserVehicleRepository.UpdateUserVehicle(v.Request.Context(), vehicleUpdate.Name, vehicleUpdate.Model, vehicleUpdate.Mileage, id)
	if err != nil {
		return nil, err
	}

	findVehicle, err = v.IUserVehicleRepository.GetUserVehicleById(v.Request.Context(), id)
	if err != nil {
		return nil, err
	}

	return findVehicle, nil
}

func (v *VehicleController) DeleteVehicle() error {
	id, err := helper.GetQueryID(v.Request)
	if err != nil {
		return err
	}
	_, err = v.IUserVehicleRepository.GetUserVehicleById(v.Request.Context(), id)
	if err != nil {
		return errors.New("vehicle do not exist in database")
	}

	if _, err = v.IUserVehicleRepository.DeleteUserVehicle(v.Request.Context(), id); err != nil {
		return errors.New("cannot delete vehicle")
	}

	return nil
}

func (v *VehicleController) GetDirections() ([]google_API.Route, error) {
	start := helper.GetQueryStart(v.Request)
	end := helper.GetQueryEnd(v.Request)
	var err error
	var clientRequest google_API.Request
	var startLat float64
	var startLng float64
	var endLat float64
	var endLng float64

	clientRequest.Origin, err = v.IShipPortsRepository.GetCityByName(v.Request.Context(), start)
	if err != nil || clientRequest.Origin == "" || clientRequest.Origin != start {
		return nil, errors.New("point of origin do not exist in database")
	}

	clientRequest.Destination, err = v.IShipPortsRepository.GetCityByName(v.Request.Context(), end)
	if err != nil || clientRequest.Destination == "" || clientRequest.Destination != end {
		return nil, errors.New("destination do not exist in database")
	}

	originLatitude, err := v.IShipPortsRepository.GetCityLatitude(v.Request.Context(), clientRequest.Origin)
	originLongitude, err := v.IShipPortsRepository.GetCityLongitude(v.Request.Context(), clientRequest.Origin)

	destinationLatitude, err := v.IShipPortsRepository.GetCityLatitude(v.Request.Context(), clientRequest.Destination)
	destinationLongitude, err := v.IShipPortsRepository.GetCityLongitude(v.Request.Context(), clientRequest.Destination)

	client := google_API.NewClientRequest(clientRequest.Origin, clientRequest.Destination)
	route, err := client.FindRoute(clientRequest)
	if err != nil {
		return nil, err
	}

	totalDistance := google_API.GetTotalDistance()

	vehicles, err := v.IUserVehicleRepository.FindVehiclesForUser(v.Request.Context(), UserID)
	mileage := maxMileage(vehicles)

	if len(vehicles) == 1 {
		if float32(totalDistance) <= vehicles[0].Mileage {
			return route, nil
		} else {
			gasStations, _ := google_API.GasStations(originLatitude, originLongitude, totalDistance)
			for i := 0; i < len(gasStations.Results); i++ {
				startLat = gasStations.Results[i].Geometry.Location.Lat
				startLng = gasStations.Results[i].Geometry.Location.Lng
				endLat = gasStations.Results[i+1].Geometry.Location.Lat
				endLng = gasStations.Results[i+1].Geometry.Location.Lng

				distanceGasStations, _ := google_API.DistanceMatrix(float32(startLat), float32(startLng), float32(endLat), float32(endLng))
				for j := 0; j < len(distanceGasStations.Rows); j++ {
					for k := 0; k < len(distanceGasStations.Rows[i].Elements); k++ {
						if distanceGasStations.Rows[i].Elements[k].Distance.Value/1000 > int(mileage) || gasStations.Results[i].OpeningHours.OpenNow == false {
							return nil, errors.New("route not possible vehicle can not reach the gas station for refueling")
						}
					}
				}
			}
			startLat = gasStations.Results[len(gasStations.Results)-1].Geometry.Location.Lat
			startLng = gasStations.Results[len(gasStations.Results)-1].Geometry.Location.Lng
			distanceGasStations, _ := google_API.DistanceMatrix(float32(startLat), float32(startLng), destinationLatitude, destinationLongitude)
			if distanceGasStations.Rows[0].Elements[0].Distance.Value/1000 > int(mileage) || gasStations.Results[0].OpeningHours.OpenNow == false {
				return nil, errors.New("route not possible vehicle can not reach the gas station for refueling")
			}
		}
	}
	return route, nil
}

func maxMileage(vehicles []*models.Vehicle) float32 {
	max := vehicles[0].Mileage

	for _, m := range vehicles {
		if max < m.Mileage {
			max = m.Mileage
		}
	}

	return max
}
