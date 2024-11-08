package repository

import (
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"context"
	"errors"
	"html"
	"strings"

	"github.com/google/uuid"
)

func NewVehicleRepo() interfaces.IUserVehicleRepository {
	return &PG{DB: GetDb()}
}

func NewVehicleObject(id uuid.UUID, name string, model string, mileage float32, userId uuid.UUID) models.Vehicle {
	userVehicle := models.Vehicle{}
	name = html.EscapeString(strings.TrimSpace(name))
	model = html.EscapeString(strings.TrimSpace(model))

	userVehicle = models.Vehicle{
		ID:      id,
		Name:    name,
		Model:   model,
		Mileage: mileage,
		UserId:  userId,
	}

	return userVehicle
}

func (u *PG) CreateUserVehicle(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error) {

	if vehicle == nil {
		return &models.Vehicle{}, errors.New("vehicle object can not be empty")
	}

	if err := u.DB.WithContext(ctx).Model(&vehicle).Create(&vehicle).Error; err != nil {
		return nil, err
	}

	return vehicle, nil

}

func (u *PG) UpdateUserVehicle(ctx context.Context, vehicleName string, vehicleModel string, vehicleMileage float32, userid uuid.UUID) (*models.Vehicle, error) {
	vehicle := &models.Vehicle{}
	err := u.DB.WithContext(ctx).Model(&models.Vehicle{}).
		Where("id = ?", userid).Take(&vehicle).
		UpdateColumns(map[string]interface{}{"name": vehicleName, "model": vehicleModel, "mileage": vehicleMileage}).Error
	if err != nil {
		err.Error()
		return nil, err
	}

	return vehicle, nil
}

func (u *PG) GetUserVehicleById(ctx context.Context, id uuid.UUID) (*models.Vehicle, error) {

	vehicle := &models.Vehicle{}

	if err := u.DB.WithContext(ctx).Model(vehicle).Where("id = ?", id).Find(vehicle).Error; err != nil {
		return &models.Vehicle{}, err
	}

	return vehicle, nil
}

func (u *PG) DeleteUserVehicle(ctx context.Context, id uuid.UUID) (int64, error) {
	var err error

	tx := u.DB.Begin()

	delTx := tx.WithContext(ctx).Model(&models.Vehicle{}).Where("id = ?", id).Delete(&models.Vehicle{})

	if err = delTx.Error; err != nil {
		return 0, err
	} else {
		tx.Commit()
	}

	return delTx.RowsAffected, nil
}

func (u *PG) FindUserVehicle(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	var err error
	user := &models.User{}

	if err = u.DB.WithContext(ctx).Joins("join vehicles ON vehicles.user_id = users.id").Find(&user).Error; err != nil {
		return nil, err
	}

	var vehicles []models.Vehicle

	if err = u.DB.WithContext(ctx).Model(&vehicles).Where("user_id = ?", userID).Find(&vehicles).Error; err != nil {
		return nil, err
	}
	user.UserVehicle = vehicles

	return user, nil
}

func (u *PG) FindVehiclesForUser(ctx context.Context, userID uuid.UUID) ([]*models.Vehicle, error) {
	var err error
	var vehicles []*models.Vehicle
	if err = u.DB.WithContext(ctx).Model(&vehicles).Where("user_id = ?", userID).Find(&vehicles).Error; err != nil {
		return nil, err
	}

	return vehicles, nil
}
