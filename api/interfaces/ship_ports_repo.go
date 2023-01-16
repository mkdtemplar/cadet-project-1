package interfaces

import (
	"cadet-project/models"
	"context"
)

type IShipPortsRepository interface {
	FindUserPrefPorts(ctx context.Context, usrpref *models.UserPreferences) (*models.UserPreferences, error)
}
