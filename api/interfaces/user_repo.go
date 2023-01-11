package interfaces

import (
	"cadet-project/models"
	"context"

	"github.com/google/uuid"
)

type IUserRepository interface {
	PrepareUserData(email string, name string)
	SaveUserDb(ctx context.Context, usr *models.User) (*models.User, error)
	DeleteUserDb(ctx context.Context, uid uuid.UUID) (int64, error)
	GetUser(ctx context.Context, user *models.User) (*models.User, error)
}
