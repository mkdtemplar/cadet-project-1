package repository

import (
	"cadet-project/models"
	"cadet-project/repository/interfaces"
	"context"
	"errors"

	"gorm.io/gorm"
)

func NewUserPrefRepo(db *gorm.DB) interfaces.IUserPreferences {
	return &PG{db: db}
}

func (u *PG) SaveUserPreferences(ctx context.Context, in *models.UserPreferences) error {
	if in == nil {
		return errors.New("user details empty")
	}

	return u.db.Debug().WithContext(ctx).Create(&in).Error
}

func (u *PG) FindUserPreferences(ctx context.Context, id uint32) (*models.UserPreferences, error) {
	var err error
	err = u.db.Debug().WithContext(ctx).Model(&models.UserPreferences{}).Where("user_id = ?", id).Take(&u.userpref).Error
	if err != nil {
		return &models.UserPreferences{}, err
	}

	return u.userpref, nil
}

func (u *PG) UpdateUserPref(ctx context.Context, userid uint32, country string) (*models.UserPreferences, error) {
	var err error

	err = u.db.Debug().WithContext(ctx).Model(&models.UserPreferences{}).Where("user_id = ?", userid).Update("country", country).Error
	if err != nil {

		err.Error()
	}

	return u.userpref, nil

}

func (u *PG) DeleteUserPreferences(ctx context.Context, usrpref *models.UserPreferences) (int64, error) {
	var err error
	userPref := &models.UserPreferences{UserId: usrpref.UserId}
	tx := u.db.Begin()

	delTx := tx.WithContext(ctx).Model(&userPref).Where("user_id = ?", usrpref.UserId).Delete(userPref)

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
	if err = u.db.WithContext(ctx).Joins("join ships_routes ON ships_routes.country = user_preferences.country").Find(userPref).Error; err != nil {
		return nil, err
	}

	var ports []models.ShipsRoutes

	if err := u.db.WithContext(ctx).Where("country = ?", in.Country).Model(&models.ShipsRoutes{}).Find(&ports).Error; err != nil {
		return &models.UserPreferencesPorts{}, err
	}
	userPrefPorts := &models.UserPreferencesPorts{
		Country: userPref.Country,
		UserId:  userPref.UserId,
		Ports:   ports,
	}

	return userPrefPorts, nil
}
