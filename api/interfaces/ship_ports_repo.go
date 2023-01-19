package interfaces

import (
	"cadet-project/models"
	"context"
)

type IShipPortsRepository interface {
	FindUserPrefPorts(ctx context.Context, usrpref *models.UserPreferences) (*models.UserPreferences, error)
	//FindPreferences(ctx context.Context, id uuid.UUID) (*models.UserPreferences, error)
	//GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
}
