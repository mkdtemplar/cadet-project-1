package interfaces

import (
	"cadet-project/pkg/models"
	"context"

	"github.com/google/uuid"
)

type IShipPortsRepository interface {
	FindUserPrefPorts(ctx context.Context, usrpref *models.UserPreferences) (*models.UserPreferences, error)
	FindUserPorts(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetCityByName(ctx context.Context, name string) (string, error)
	GetCityLatitude(ctx context.Context, name string) (float32, error)
	GetCityLongitude(ctx context.Context, name string) (float32, error)
}
