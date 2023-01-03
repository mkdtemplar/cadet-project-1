package interfaces

import (
	"cadet-project/models"
	"context"

	"github.com/google/uuid"
)

type IUserRepository interface {
	PrepareUserData(email string, name string)
	SaveUserDb(ctx context.Context, create *models.User) error
	DeleteUserDb(ctx context.Context, uid uuid.UUID) (int64, error)
	GetUser(ctx context.Context, user *models.User) (*models.User, error)
}
