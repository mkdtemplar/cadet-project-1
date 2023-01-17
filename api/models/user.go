package models

import (
	"cadet-project/utils"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID         `gorm:"primary_key;type:uuid" json:"id"`
	Email    string            `gorm:"size:100;not null;unique" json:"email"`
	Name     string            `gorm:"size:100" json:"name"`
	UserPref []UserPreferences `gorm:"foreignKey:UserId;references:ID"`
}

func (u *User) Clean() {
	u.Email = utils.Clean(u.Email)
	u.Name = utils.Clean(u.Name)
}
