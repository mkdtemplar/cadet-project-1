package repository

import (
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"cadet-project/pkg/repository/generate_id"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PG struct {
	DB *gorm.DB
}

func NewUserRepo() interfaces.IUserRepository {
	return &PG{DB: GetDb()}
}

func (u *PG) Create(ctx context.Context, usr *models.User) (*models.User, error) {
	if usr == nil {
		return &models.User{}, errors.New("user details empty")
	}
	usr.ID = generate_id.GenerateID()
	err := u.DB.WithContext(ctx).Model(models.User{}).Create(&usr).Error

	if err != nil {
		return &models.User{}, err
	}

	return usr, nil
}

func (u *PG) Delete(ctx context.Context, uid uuid.UUID) (int64, error) {
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

func (u *PG) GetUserEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}

	err := u.DB.WithContext(ctx).Model(&models.User{}).Where("email= ?", email).Find(&user).Error

	userFind := &models.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	return userFind, err
}

func (u *PG) GetById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var err error

	user := &models.User{}

	err = u.DB.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	userFind := &models.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	return userFind, nil
}
