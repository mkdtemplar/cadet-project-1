package interfaces

import (
	"cadet-project/models"
	"context"

	"github.com/google/uuid"
)

type IUserRepository interface {
	PrepareUserData(email string, name string)
	Create(ctx context.Context, usr *models.User) (*models.User, error)
	Delete(ctx context.Context, uid uuid.UUID) (int64, error)
	GetUserEmail(ctx context.Context, email string) (*models.User, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.User, error)
}
