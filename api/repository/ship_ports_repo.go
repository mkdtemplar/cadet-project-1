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
	userPref := &models.UserPreferences{}
	if err = u.DB.WithContext(ctx).Joins("join ship_ports ON ship_ports.country = user_preferences.user_country").Find(userPref).Error; err != nil {
		return nil, err
	}

	var ports []models.ShipPorts

	if err := u.DB.WithContext(ctx).Where("country = ?", in.UserCountry).Model(&models.ShipPorts{}).Find(&ports).Error; err != nil {
		return &models.UserPreferences{}, err
	}

	userPref.Ports = ports

	return userPref, nil
}
