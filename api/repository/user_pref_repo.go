package repository

import (
	"cadet-project/interfaces"
	"cadet-project/models"
	"cadet-project/repository/generate_id"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewUserPrefRepo(db *gorm.DB) interfaces.IUserPreferencesRepository {
	return &PG{DB: db}
}

func (u *PG) SaveUserPreferences(ctx context.Context, in *models.UserPreferences) error {
	if in == nil {
		return errors.New("user details empty")
	}
	in.ID = generate_id.GenerateID()
	return u.DB.WithContext(ctx).Create(&in).Error
}

func (u *PG) FindUserPreferences(ctx context.Context, id uuid.UUID) (*models.UserPreferences, error) {
	var err error

	userPref := &models.UserPreferences{}

	err = u.DB.WithContext(ctx).Model(&models.UserPreferences{}).Where("user_id = ?", id).Take(userPref).Error
	if err != nil {
		return nil, nil
	}

	return userPref, nil
}

func (u *PG) UpdateUserPref(ctx context.Context, userid uuid.UUID, country string) (*models.UserPreferences, error) {
	var err error

	err = u.DB.WithContext(ctx).Model(&models.UserPreferences{}).Where("user_id = ?", userid).Update("country", country).Error
	if err != nil {
		err.Error()
		return nil, err
	}

	return u.UserPreferences, nil

}

func (u *PG) DeleteUserPreferences(ctx context.Context, id uuid.UUID) (int64, error) {
	var err error

	tx := u.DB.Begin()

	delTx := tx.WithContext(ctx).Model(&models.UserPreferences{}).Where("user_id = ?", id).Delete(&models.User{}, id)

	if err = delTx.Error; err != nil {
		return 0, err
	} else {
		tx.Commit()
	}

	return delTx.RowsAffected, nil
}

func (u *PG) FindUserPrefPorts(ctx context.Context, in *models.UserPreferences) (*models.UserPreferencesPorts, error) {
	var err error
	userPref := &models.UserPreferences{}
	if err = u.DB.WithContext(ctx).Joins("join ships_routes ON ships_routes.country = user_preferences.country").Find(userPref).Error; err != nil {
		return nil, err
	}

	var ports []models.ShipsRoutes

	if err := u.DB.WithContext(ctx).Where("country = ?", in.Country).Model(&models.ShipsRoutes{}).Find(&ports).Error; err != nil {
		return &models.UserPreferencesPorts{}, err
	}
	userPrefPorts := &models.UserPreferencesPorts{
		Country: userPref.Country,
		UserId:  userPref.UserId,
		Ports:   ports,
	}

	return userPrefPorts, nil
}
