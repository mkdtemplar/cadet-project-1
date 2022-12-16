package models

import (
	"cadet-project/responses"
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"html"
	"net/http"
	"strings"
	"time"
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
	tokenName, err := r.Cookie("token")

	if err != nil {
		return ""
	}
	return tokenName.Value
}

func CreateCookieToAllEndPoints(tokenValue string, exp time.Time) http.Cookie {

	cookie := http.Cookie{
		Name:     "session_token",
		Value:    tokenValue,
		HttpOnly: false,
		Path:     "/",
		Expires:  exp,
	}

	return cookie
}

func TokenValid(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("token")
	if err != nil {
		return err
	}
	sessionToken := cookie.Value
	userSession, exists := Sessions[sessionToken]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("token not present in session"))
		return errors.New("invalid token")
	}

	if userSession.IsExpired() {
		delete(Sessions, sessionToken)
		responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return errors.New("unauthorized")
	}

	return nil
}
