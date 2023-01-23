package interfaces

import (
	models2 "cadet-project/pkg/models"
	"context"

	"github.com/google/uuid"
)

type IShipPortsRepository interface {
	FindUserPrefPorts(ctx context.Context, usrpref *models2.UserPreferences) (*models2.UserPreferences, error)
	FindUserPorts(ctx context.Context, id uuid.UUID) (*models2.User, error)
}
