package models

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `gorm:"primary_key" json:"id"`
	Email string    `gorm:"size:100;not null;unique" json:"email"`
	Name  string    `gorm:"size:100" json:"name"`
}
