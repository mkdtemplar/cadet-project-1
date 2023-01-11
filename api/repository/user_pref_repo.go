package repository

import (
	"cadet-project/interfaces"
	"cadet-project/models"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewUserPrefRepo(db *gorm.DB) interfaces.IUserPreferencesRepository {
	return &PG{DB: db}
}

func (u *PG) GetAllUserPreferences(ctx context.Context, userId uuid.UUID) ([]*models.UserPreferences, error) {
	var userPrefList []*models.UserPreferences

	err := u.DB.WithContext(ctx).Model(&models.UserPreferences{}).Where("user_id = ?", userId).Find(&userPrefList).Error
	if err != nil {
		return nil, err
	}

	return userPrefList, nil
}

func (u *PG) SaveUserPreferences(ctx context.Context, in *models.UserPreferences) (*models.UserPreferences, error) {
	if in == nil {
		return &models.UserPreferences{}, errors.New("user details empty")
	}

	err := u.DB.WithContext(ctx).Model(u.UserPreferences).Create(&in).Error
	if err != nil {
		return &models.UserPreferences{}, errors.New("user not created")
	}
	return in, nil
}

func (u *PG) FindUserPreferences(ctx context.Context, id uuid.UUID) (*models.UserPreferences, error) {
	var err error

	userPref := &models.UserPreferences{}

	err = u.DB.WithContext(ctx).Model(&models.UserPreferences{}).Where("id = ?", id).Find(&userPref).Error
	if err != nil {
		return nil, err
	}

	return userPref, nil
}

func (u *PG) UpdateUserPref(ctx context.Context, userid uuid.UUID, country string) (*models.UserPreferences, error) {
	var err error

	err = u.DB.WithContext(ctx).Model(&models.UserPreferences{}).Where("id = ?", userid).Update("country", country).Error
	if err != nil {
		err.Error()
		return nil, err
	}

	return u.UserPreferences, nil

}

func (u *PG) DeleteUserPreferences(ctx context.Context, id uuid.UUID) (int64, error) {
	var err error

	tx := u.DB.Begin()

	delTx := tx.WithContext(ctx).Model(&models.UserPreferences{}).Where("id = ?", id).Delete(&models.UserPreferences{})

	if err = delTx.Error; err != nil {
		return 0, err
	} else {
		tx.Commit()
	}

	return delTx.RowsAffected, nil
}

func (u *PG) FindUserPrefPorts(ctx context.Context, in *models.UserPreferences) (*models.UserPreferences, error) {
	var err error
	userPref := &models.UserPreferences{}
	if err = u.DB.WithContext(ctx).Joins("join ships_routes ON ships_routes.country = user_preferences.user_country").Find(userPref).Error; err != nil {
		return nil, err
	}

	var ports []models.ShipsRoutes

	if err := u.DB.WithContext(ctx).Where("country = ?", in.UserCountry).Model(&models.ShipsRoutes{}).Find(&ports).Error; err != nil {
		return &models.UserPreferences{}, err
	}

	userPref.Ports = ports

	return userPref, nil
}
