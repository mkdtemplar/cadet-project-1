package models

import "github.com/google/uuid"

type UserPreferences struct {
	ID          uuid.UUID   `gorm:"primary_key;type:uuid" json:"id"`
	UserCountry string      `json:"user_country"`
	UserId      uuid.UUID   `gorm:"type:uuid" json:"user_id"`
	Ports       []ShipPorts `gorm:"-" json:"ports"`
}
