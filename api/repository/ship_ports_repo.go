package repository

import (
	"cadet-project/interfaces"
	"cadet-project/models"
	"context"

	"github.com/google/uuid"
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

	return &userPref, nil
}

func (u *PG) FindPreferences(ctx context.Context, id uuid.UUID) (*models.UserPreferences, error) {
	var err error

	userPref := models.UserPreferences{}

	err = u.DB.WithContext(ctx).Model(&models.UserPreferences{}).Where("id = ?", id).Take(&userPref).Error
	if err != nil {
		return &models.UserPreferences{}, err
	}

	return &userPref, nil
}
