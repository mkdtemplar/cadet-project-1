package interfaces

import (
	"cadet-project/models"
	"context"
)

type IUserPreferences interface {
	SaveUserPreferences(ctx context.Context, usrpref *models.UserPreferences) error
	FindUserPreferences(ctx context.Context, id uint32) (*models.UserPreferences, error)
	UpdateUserPref(ctx context.Context, id uint32, country string) (*models.UserPreferences, error)
	DeleteUserPreferences(ctx context.Context, usrpref *models.UserPreferences) (int64, error)
	FindUserPrefPorts(ctx context.Context, usrpref *models.UserPreferences) (*models.UserPreferencesPorts, error)
}
