package repository

import (
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/repository/generate_id"
	"context"
	"errors"
	"html"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewVehicleRepo(db *gorm.DB) interfaces.IUserVehicleRepository {
	return &PG{DB: db}
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
		return vehicle, errors.New("vehicle object can not be empty")
	}

	userId, _ := uuid.Parse("26810825-80e2-46be-b72b-744c2de4a872")

	vehicle = &models.Vehicle{
		ID:      generate_id.GenerateID(),
		Name:    "Audi",
		Model:   "R8",
		Mileage: 200,
		UserId:  userId,
	}

	if err := u.DB.WithContext(ctx).Model(vehicle).Create(&vehicle).Error; err != nil {
		return nil, err
	}

	return vehicle, nil

}

func (u *PG) UpdateUserVehicle(ctx context.Context, vehicle *models.Vehicle, userid uuid.UUID) (*models.Vehicle, error) {
	var err error
	userid, _ = uuid.Parse("26810825-80e2-46be-b72b-744c2de4a872")

	vehicle = &models.Vehicle{
		ID:      generate_id.GenerateID(),
		Name:    "Audi",
		Model:   "R8",
		Mileage: 200,
		UserId:  userid,
	}

	err = u.DB.WithContext(ctx).Model(vehicle).
		Where("id = ?", userid).Take(&vehicle).
		UpdateColumns(map[string]interface{}{"name": "Audi", "model": "QRS 8", "mileage": 400}).Error
	if err != nil {
		err.Error()
		return nil, err
	}

	vehicleUpdated := &models.Vehicle{
		Name:    "Audi",
		Model:   "QRS 8",
		Mileage: 400,
		UserId:  userid,
	}

	return vehicleUpdated, nil

}

func (u *PG) GetUserVehicleById(ctx context.Context, id uuid.UUID) (*models.Vehicle, error) {

	vehicle := &models.Vehicle{}

	userid, _ := uuid.Parse("26810825-80e2-46be-b72b-744c2de4a872")
	id, _ = uuid.Parse("af82a7bd-fb18-4c68-ab32-b10221735f2d")

	if err := u.DB.WithContext(ctx).Model(vehicle).Where("id = ?", id).Find(vehicle).Error; err != nil {
		return nil, err
	}

	vehicle = &models.Vehicle{
		ID:      id,
		Name:    "Audi",
		Model:   "R8",
		Mileage: 200,
		UserId:  userid,
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
