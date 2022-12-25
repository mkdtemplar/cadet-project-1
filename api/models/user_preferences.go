package models

import (
	"errors"
	"html"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
)

type UserPreferences struct {
	ID      uint32        `gorm:"primary_key;auto_increment" json:"id"`
	Country string        `json:"country"`
	UserId  uint32        `json:"user_id"`
	Ports   []ShipsRoutes `json:"ports"`
}

func (up *UserPreferences) ConstructUserPrefObject(country string, userid uint32) {
	country = html.EscapeString(strings.TrimSpace(country))
	userid = up.UserId
}

func (up *UserPreferences) ValidateUserPref(country string, userid uint32) error {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]+$`)
	checkNumber := regexp.MustCompile(`\d+`)

	if country == "" || checkLetters.MatchString(country) == false {
		return errors.New("country cannot be empty")
	}

	if up.UserId < 1 || checkNumber.MatchString(string(userid)) {
		return errors.New("user id is required")
	}

	return nil
}

func (up *UserPreferences) SaveUserPreferences(db *gorm.DB) (*UserPreferences, error) {
	var err error

	err = db.Debug().Model(&UserPreferences{}).Create(&up).Error
	if err != nil {
		return &UserPreferences{}, err
	}

	return up, nil
}

func (up *UserPreferences) FindUserPreferences(db *gorm.DB, id uint32) (*UserPreferences, error) {
	var err error
	err = db.Debug().Model(&UserPreferences{}).Where("id = ?", id).Take(&up).Error
	if err != nil {
		return &UserPreferences{}, err
	}

	return up, nil
}

func (up *UserPreferences) UpdateUserPref(db *gorm.DB) (*UserPreferences, error) {

	var err error

	err = db.Debug().Model(&UserPreferences{}).Where("id = ?", up.ID).Updates(UserPreferences{Country: up.Country}).Error
	if err != nil {

		err.Error()
	}

	return up, nil
}

func (up *UserPreferences) DeleteUserPreferences(db *gorm.DB, userid uint32) (int64, error) {

	db = db.Debug().Model(&UserPreferences{}).Where("id = ?", userid).Take(&UserPreferences{}).Delete(&UserPreferences{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("user preferences record not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (up *UserPreferences) FindUserPrefPorts(db *gorm.DB, country string) (*UserPreferences, error) {
	var err error

	if err = db.Joins("join ships_routes ON ships_routes.country = user_preferences.country").Find(&up).Error; err != nil {
		return nil, err
	}
	var ports []ShipsRoutes

	if err := db.Where("country = ?", country).Model(&ShipsRoutes{}).Find(&ports).Error; err != nil {
		return &UserPreferences{}, err
	}
	up.Ports = ports

	return up, nil
}

func (up *UserPreferences) FindCountry(db *gorm.DB, country string) (*string, error) {

	err := db.Debug().Model(&UserPreferences{}).Where("country = ?", country).Take(&up).Error
	if err != nil {
		return nil, err
	}
	return &up.Country, nil
}
