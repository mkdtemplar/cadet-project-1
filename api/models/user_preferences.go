package models

type UserPreferences struct {
	ID      uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Country string `json:"country"`
	UserId  uint32 `json:"user_id"`
	Name    string `json:"name"`
}
