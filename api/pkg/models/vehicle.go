package models

import "github.com/google/uuid"

type Vehicle struct {
	ID      uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	Name    string    `gorm:"size:100" json:"name"`
	Model   string    `gorm:"size:100" json:"model"`
	Mileage float32   `json:"mileage"`
	UserId  uuid.UUID `gorm:"type:uuid" json:"user_id"`
}

// separate gorm and JSON
