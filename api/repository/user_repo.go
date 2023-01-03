package repository

import (
	"cadet-project/interfaces"
	"cadet-project/models"
	"context"
	"errors"
	"html"
	"strings"

	"gorm.io/gorm"
)

type PG struct {
	DB              *gorm.DB
	User            *models.User
	UserPreferences *models.UserPreferences
}

func NewUserRepo(db *gorm.DB) interfaces.IUserRepository {
	return &PG{DB: db}
}

func (u *PG) PrepareUserData(email string, name string) {
	email = html.EscapeString(strings.TrimSpace(email))
	name = html.EscapeString(strings.TrimSpace(name))
}

func (u *PG) SaveUserDb(ctx context.Context, usr *models.User) error {
	if usr == nil {
		return errors.New("user details empty")
	}

	return u.DB.Debug().WithContext(ctx).Create(&usr).Error
}

func (u *PG) DeleteUserDb(ctx context.Context, uid uint64) (int64, error) {
	var err error

	tx := u.DB.Begin()

	delTx := tx.WithContext(ctx).Model(&models.User{}).Delete(&models.User{}, uid)

	if err = delTx.Error; err != nil {
		return 0, err
	} else {
		tx.Commit()
	}

	return tx.RowsAffected, nil
}

func (u *PG) GetUser(ctx context.Context, in *models.User) (*models.User, error) {
	user := &models.User{}

	err := u.DB.Debug().WithContext(ctx).Take(user, "email = ?", in.Email).Error
	return user, err
}
