package repository

import (
	"cadet-project/interfaces"
	"cadet-project/models"
	"context"

	"gorm.io/gorm"
)

func NewShipPortsRepo(db *gorm.DB) interfaces.IShipPortsRepository {
	return &PG{DB: db}
}

func (u *PG) FindUserPrefPorts(ctx context.Context, in *models.UserPreferences) (*models.UserPreferences, error) {
	var err error
	userPref := models.UserPreferences{}
	if err = u.DB.WithContext(ctx).Joins("join ship_ports ON ship_ports.country = user_preferences.user_country").Find(&userPref).Error; err != nil {
		return nil, err
	}

	var ports []models.ShipPorts

	if err := u.DB.WithContext(ctx).Where("country = ?", in.UserCountry).Model(&models.ShipPorts{}).Find(&ports).Error; err != nil {
		return &models.UserPreferences{}, err
	}

	userPref.Ports = ports
	userPref.UserCountry = in.UserCountry
	userPref.UserId = in.UserId

	return &userPref, nil
}

func (u *PG) FindUserPorts(ctx context.Context, usr *models.User) (*models.User, error) {
	var err error

	user := models.User{}
	if err = u.DB.WithContext(ctx).Model(&models.User{}).Joins("join user_preferences ON user_preferences.user_id = users.id").Find(&user).Error; err != nil {
		return nil, err
	}
	var userPref models.UserPreferences
	if err = u.DB.WithContext(ctx).Where("user_id = ?", usr.ID).Model(&models.UserPreferences{}).Find(&userPref).Error; err != nil {
		return nil, err
	}

	if err = u.DB.WithContext(ctx).Joins("join ship_ports ON ship_ports.country = user_preferences.user_country").Find(&userPref).Error; err != nil {
		return nil, err
	}

	var ports []models.ShipPorts

	if err := u.DB.WithContext(ctx).Where("country = ?", userPref.UserCountry).Model(&models.ShipPorts{}).Find(&ports).Error; err != nil {
		return nil, err
	}

	userPref.Ports = ports
	user.UserPref = userPref
	return &user, nil
}
