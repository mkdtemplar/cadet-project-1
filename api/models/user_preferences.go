package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type UserPreferences struct {
	ID      uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Country string `json:"country"`
	UserId  uint32 `json:"user_id"`
	Port    []Ports
}

func (up *UserPreferences) PrepareUserPref() {
	up.ID = 0
	up.Country = html.EscapeString(strings.TrimSpace(up.Country))
	up.UserId = 0
}

func (up *UserPreferences) ValidateUserPref() error {
	if up.Country == "" {
		return errors.New("country cannot be empty")
	}

	if up.UserId < 1 {
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

func (up *UserPreferences) FindAllUserPref(db *gorm.DB) (*[]UserPreferences, error) {
	var err error
	var userPref []UserPreferences
	err = db.Debug().Model(&UserPreferences{}).Limit(100).Find(&userPref).Error
	if err != nil {
		return &[]UserPreferences{}, err
	}

	return &userPref, nil
}

func (up *UserPreferences) FindOneUserPref(db *gorm.DB, id uint32) (*UserPreferences, error) {
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
		return &UserPreferences{}, err
	}

	return up, nil
}

func (up *UserPreferences) DeleteUserPref(db *gorm.DB, userid uint32) (int64, error) {

	db = db.Debug().Model(&UserPreferences{}).Where("id = ?", userid).Take(&UserPreferences{}).Delete(&UserPreferences{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("user preferences record not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

/*
func (up *UserPreferences) FindUserPrefPorts(db *gorm.DB, country string) (*UserPreferences, error) {
	var err error

	err = db.Debug().Model(&UserPreferences{}).Where("country =?", country).Take(&up.Country).Error
	if err != nil {
		return &UserPreferences{}, err
	}

	if up.Country == country {
		return up, errors.New("user preferences country not found")
	}
	var ports []Ports

	err = db.Raw("")
}


*/
