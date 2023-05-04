package models

import (
	"cadet-project/pkg/utils"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID       `gorm:"primary_key;type:uuid" json:"id"`
	Email       string          `gorm:"size:100;not null;unique" json:"email"`
	Name        string          `gorm:"size:100" json:"name"`
	UserPref    UserPreferences `gorm:"foreignKey:UserId;references:ID" json:"user_pref"`
	UserVehicle []Vehicle       `gorm:"foreignKey:UserId;references:ID" json:"user_vehicle"`
}

func (u *User) Clean() {
	u.Email = utils.CleanUserData(u.Email)
	u.Name = utils.CleanUserData(u.Name)
}
