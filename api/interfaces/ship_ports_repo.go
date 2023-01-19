package interfaces

import (
	"cadet-project/models"
	"context"

	"github.com/google/uuid"
)

type IShipPortsRepository interface {
	FindUserPrefPorts(ctx context.Context, usrpref *models.UserPreferences) (*models.UserPreferences, error)
	FindUserPorts(ctx context.Context, id uuid.UUID) (*models.User, error)
}
