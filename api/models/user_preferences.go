package models

type UserPreferences struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Country  string `json:"country"`
	UseridFk uint32 `json:"user_id_fk"`
	Name     string `json:"name"`
}
