package repository

import (
	"cadet-project/interfaces"
	"cadet-project/models"
	"cadet-project/repository/generate_id"
	"context"
	"errors"
	"html"
	"strings"

	"github.com/google/uuid"
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

func (u *PG) SaveUserDb(ctx context.Context, usr *models.User) (*models.User, error) {
	if usr == nil {
		return &models.User{}, errors.New("user details empty")
	}
	usr.ID = generate_id.GenerateID()
	err := u.DB.WithContext(ctx).Model(u.User).Create(&usr).Error

	if err != nil {
		return &models.User{}, err
	}

	return usr, nil
}

func (u *PG) DeleteUserDb(ctx context.Context, uid uuid.UUID) (int64, error) {
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

	err := u.DB.WithContext(ctx).Preload("UserPref").Take(user, "email = ?", in.Email).Error
	return user, err
}
