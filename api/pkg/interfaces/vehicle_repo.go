package interfaces

import (
	"cadet-project/pkg/models"
	"context"

	"github.com/google/uuid"
)

type IUserVehicleRepository interface {
	CreateUserVehicle(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error)
	UpdateUserVehicle(ctx context.Context, vehicleName string, vehicleModel string, vehicleMileage float32, userid uuid.UUID) (*models.Vehicle, error)
	GetUserVehicleById(ctx context.Context, id uuid.UUID) (*models.Vehicle, error)
	DeleteUserVehicle(ctx context.Context, id uuid.UUID) (int64, error)
	FindUserVehicle(ctx context.Context, userID uuid.UUID) (*models.User, error)
	FindVehiclesForUser(ctx context.Context, userID uuid.UUID) ([]*models.Vehicle, error)
	FindVehicleWithUserID(ctx context.Context, userID uuid.UUID) ([]*models.Vehicle, error)
}
