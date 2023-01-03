package interfaces

import (
	"cadet-project/models"
	"context"
)

type IUserRepository interface {
	PrepareUserData(email string, name string)
	SaveUserDb(ctx context.Context, usr *models.User) error
	DeleteUserDb(ctx context.Context, uid uint64) (int64, error)
	CheckUser(ctx context.Context, user *models.User) (*models.User, error)
}
