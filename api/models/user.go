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
}

func (u *User) PrepareUserData() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) ValidateUserData(action string) (err error) {
	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			return errors.New("e-mail is required")
		}
		if err = checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid E-mail format")
		}
	default:
		if u.Email == "" {
			return errors.New("e-mail is required")
		}
		if err = checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid E-mail format")
		}
	}
	return
}

func (u *User) SaveUserDb(db *gorm.DB) error {
	return db.Debug().Create(u).Error
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
		HttpOnly: true,
	}

	return cookie
}

func TokenValid(r *http.Request) error {
	tokenValue := SetCookieToAllEndPoints(r)

	if tokenValue.Valid() != nil {
		return errors.New("token expired")
	}

	return nil
}