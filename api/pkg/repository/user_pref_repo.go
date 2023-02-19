package repository

import (
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"context"
	"errors"
	"html"
	"strings"

	"github.com/google/uuid"
)

func NewUserPrefObject(id uuid.UUID, country string, userId uuid.UUID) models.UserPreferences {
	userPref := models.UserPreferences{}
	country = html.EscapeString(strings.TrimSpace(country))

	userPref = models.UserPreferences{
		ID:          id,
		UserCountry: country,
		UserId:      userId,
	}
	return userPref
}

func NewUserPrefRepo() interfaces.IUserPreferencesRepository {
	return &PG{DB: GetDb()}
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

	err := u.DB.WithContext(ctx).Model(models.UserPreferences{}).Create(&in).Error
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
	userPref := &models.UserPreferences{}
	err = u.DB.WithContext(ctx).Model(&models.UserPreferences{}).
		Where("id = ?", userid).Take(&userPref).UpdateColumns(map[string]interface{}{"user_country": country}).Error
	if err != nil {
		err.Error()
		return nil, err
	}

	return userPref, nil

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
