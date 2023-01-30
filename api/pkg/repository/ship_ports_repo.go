package repository

import (
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/models"
	"errors"

	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewShipPortsRepo(db *gorm.DB) interfaces.IShipPortsRepository {
	return &PG{DB: db}
}

func (u *PG) FindUserPrefPorts(ctx context.Context, in *models.UserPreferences) (*models.UserPreferences, error) {
	var err error
	userPref := &models.UserPreferences{}
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

	return userPref, nil
}

func (u *PG) FindUserPorts(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var err error

	user := &models.User{}
	if err = u.DB.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Find(&user).Error; err != nil {
		return &models.User{}, errors.New("user not found")
	}

	userPref := models.UserPreferences{}
	if err = u.DB.WithContext(ctx).Model(&models.UserPreferences{}).Where("user_id = ?", id).Find(&userPref).Error; err != nil {
		return &models.User{}, err
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
	return user, nil
}
