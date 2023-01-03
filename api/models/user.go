package models

type User struct {
	ID    uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Email string `gorm:"size:100;not null;unique" json:"email"`
	Name  string `gorm:"size:100" json:"name"`
}
