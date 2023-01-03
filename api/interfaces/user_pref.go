package interfaces

import (
	"cadet-project/models"
	"context"

	"github.com/google/uuid"
)

type IUserPreferences interface {
	SaveUserPreferences(ctx context.Context, usrpref *models.UserPreferences) error
	FindUserPreferences(ctx context.Context, id uuid.UUID) (*models.UserPreferences, error)
	UpdateUserPref(ctx context.Context, id uuid.UUID, country string) (*models.UserPreferences, error)
	DeleteUserPreferences(ctx context.Context, id uuid.UUID) (int64, error)
	FindUserPrefPorts(ctx context.Context, usrpref *models.UserPreferences) (*models.UserPreferencesPorts, error)
}
