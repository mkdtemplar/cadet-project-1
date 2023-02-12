package repository

import (
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewVehicleRepo(db *gorm.DB) interfaces.IUserVehicleRepository {
	return &PG{DB: db}
}

func (u *PG) CreateUserVehicle(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error) {
	//TODO implement me
	panic("implement me")
}

func (u *PG) UpdateUserVehicle(ctx context.Context, vehicle *models.Vehicle, id uuid.UUID) (*models.Vehicle, error) {
	//TODO implement me
	panic("implement me")
}

func (u *PG) GetUserVehicleById(ctx context.Context, id uuid.UUID) (*models.Vehicle, error) {
	//TODO implement me
	panic("implement me")
}

func (u *PG) DeleteUserVehicle(ctx context.Context, id uuid.UUID) (int64, error) {
	//TODO implement me
	panic("implement me")
}
