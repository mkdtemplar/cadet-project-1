package models

type User struct {
	ID    uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Email string `json:"email" gorm:"unique"`
}
