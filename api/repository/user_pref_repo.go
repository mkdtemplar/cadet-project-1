package repository

import (
	"cadet-project/interfaces"
	"cadet-project/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

func NewUserPrefRepo(db *gorm.DB) interfaces.IUserPreferences {
	return &PG{DB: db}
}

func (u *PG) SaveUserPreferences(ctx context.Context, in *models.UserPreferences) error {
	if in == nil {
		return errors.New("User details empty")
	}

	return u.DB.Debug().WithContext(ctx).Create(&in).Error
}

func (u *PG) FindUserPreferences(ctx context.Context, id uint32) (*models.UserPreferences, error) {
	var err error
	err = u.DB.Debug().WithContext(ctx).Model(&models.UserPreferences{}).Where("user_id = ?", id).Take(&u.UserPreferences).Error
	if err != nil {
		return &models.UserPreferences{}, err
	}

	return u.UserPreferences, nil
}

func (u *PG) UpdateUserPref(ctx context.Context, userid uint32, country string) (*models.UserPreferences, error) {
	var err error

	err = u.DB.Debug().WithContext(ctx).Model(&models.UserPreferences{}).Where("user_id = ?", userid).Update("country", country).Error
	if err != nil {

		err.Error()
	}

	return u.UserPreferences, nil

}

func (u *PG) DeleteUserPreferences(ctx context.Context, id uint32) (int64, error) {
	var err error

	tx := u.DB.Begin()

	delTx := tx.WithContext(ctx).Model(&models.UserPreferences{}).Delete(&models.User{}, id)

	if err = delTx.Error; err != nil {
		return 0, err
	} else {
		tx.Commit()
	}

	return tx.RowsAffected, nil
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
