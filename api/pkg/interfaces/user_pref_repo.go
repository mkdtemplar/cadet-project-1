package interfaces

import (
	"cadet-project/pkg/models"
	"context"

	"github.com/google/uuid"
)

type IUserPreferencesRepository interface {
	SaveUserPreferences(ctx context.Context, usrpref *models.UserPreferences) (*models.UserPreferences, error)
	FindUserPreferences(ctx context.Context, id uuid.UUID) (*models.UserPreferences, error)
	UpdateUserPref(ctx context.Context, id uuid.UUID, country string) (*models.UserPreferences, error)
	DeleteUserPreferences(ctx context.Context, id uuid.UUID) (int64, error)
}
