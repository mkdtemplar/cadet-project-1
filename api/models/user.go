package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"html"
	"net/http"
	"strings"
)

type User struct {
	ID    uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Email string `gorm:"size:100;not null;unique" json:"email"`
	Name  string `gorm:"size:100" json:"name"`
}

func (u *User) PrepareUserData() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) ValidateUserData(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			return errors.New("e-mail is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid E-mail format")
		}
		return nil
	default:
		if u.Email == "" {
			return errors.New("e-mail is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid E-mail format")
		}
		return nil
	}

}

func (u *User) SaveUserDb(db *gorm.DB) (*User, error) {
	var err error

	err = db.Debug().Create(&u).Error
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

func ExtractToken(r *http.Request) string {
	var err error
	var tokenName *http.Cookie
	tokenName, err = r.Cookie("token")
	if err != nil {
		return ""
	}

	return tokenName.Value
}

func SetCookieToAllEndPoints(r *http.Request) http.Cookie {
	tokenValue := ExtractToken(r)

	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenValue,
		MaxAge:   5 * 60,
		HttpOnly: false,
		Path:     "/userpref",
	}

	return cookie
}

func TokenValid(r *http.Request) error {
	tokenValue := SetCookieToAllEndPoints(r)
	/*tokenValue, err := r.Cookie("token")
	if err != nil {
		return errors.New("Invalid token ")
	}

	*/

	if tokenValue.Valid() != nil {
		return errors.New("token expired")
	}

	return nil
}
