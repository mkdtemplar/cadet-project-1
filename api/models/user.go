package models

import (
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID    uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Email string `gorm:"size:100;not null;unique" json:"email"`
	Name  string `gorm:"size:100" json:"name"`
}

func (u *User) PrepareUserData(email string, name string) {

	email = html.EscapeString(strings.TrimSpace(email))
	email = u.Email
	name = html.EscapeString(strings.TrimSpace(name))
	name = u.Name
}

func (u *User) SaveUserDb(db *gorm.DB) (*User, error) {

	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) DeleteUserDb(db *gorm.DB, uid uint64) (int64, error) {
	var err error

	tx := db.Begin()

	delTx := tx.Delete(&User{}, uid)

	if err = delTx.Error; err != nil {
		return 0, err
	} else {
		tx.Commit()
	}

	return tx.RowsAffected, nil
}

func (u *User) CheckUser(db *gorm.DB, mail string) error {
	err := db.Debug().Model(&User{}).Where("email = ?", mail).Find(&u).Error
	return err
}
