package models

import "github.com/google/uuid"

type UserPreferences struct {
	ID      uuid.UUID `gorm:"primary_key" json:"id"`
	Country string    `json:"country"`
	UserId  uuid.UUID `json:"user_id"`
}

type UserPreferencesPorts struct {
	Country string        `json:"country"`
	UserId  uuid.UUID     `json:"user_id"`
	Ports   []ShipsRoutes `json:"ports"`
}
