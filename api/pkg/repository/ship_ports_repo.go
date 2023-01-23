package repository

import (
	"cadet-project/pkg/interfaces"
	models2 "cadet-project/pkg/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewShipPortsRepo(db *gorm.DB) interfaces.IShipPortsRepository {
	return &PG{DB: db}
}

func (u *PG) FindUserPrefPorts(ctx context.Context, in *models2.UserPreferences) (*models2.UserPreferences, error) {
	var err error
	userPref := &models2.UserPreferences{}
	if err = u.DB.WithContext(ctx).Joins("join ship_ports ON ship_ports.country = user_preferences.user_country").Find(&userPref).Error; err != nil {
		return nil, err
	}

	var ports []models2.ShipPorts

	if err := u.DB.WithContext(ctx).Where("country = ?", in.UserCountry).Model(&models2.ShipPorts{}).Find(&ports).Error; err != nil {
		return &models2.UserPreferences{}, err
	}

	userPref.Ports = ports
	userPref.UserCountry = in.UserCountry
	userPref.UserId = in.UserId

	return userPref, nil
}

func (u *PG) FindUserPorts(ctx context.Context, id uuid.UUID) (*models2.User, error) {
	var err error

	user := &models2.User{}
	if err = u.DB.WithContext(ctx).Model(&models2.User{}).Where("id = ?", id).Find(&user).Error; err != nil {
		return &models2.User{}, err
	}

	userPref := models2.UserPreferences{}
	if err = u.DB.WithContext(ctx).Model(&models2.UserPreferences{}).Where("user_id = ?", id).Find(&userPref).Error; err != nil {
		return &models2.User{}, err
	}

	if err = u.DB.WithContext(ctx).Joins("join ship_ports ON ship_ports.country = user_preferences.user_country").Find(&userPref).Error; err != nil {
		return nil, err
	}

	var ports []models2.ShipPorts

	if err := u.DB.WithContext(ctx).Where("country = ?", userPref.UserCountry).Model(&models2.ShipPorts{}).Find(&ports).Error; err != nil {
		return nil, err
	}

	userPref.Ports = ports
	user.UserPref = userPref
	return user, nil
}
