package models

import (
	"cadet-project/responses"
	"errors"
	"html"
	"net/http"
	"net/mail"
	"regexp"
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
	name = html.EscapeString(strings.TrimSpace(name))
}

func ValidateUserData(action string, email string, name string) error {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]*$`)
	name = strings.ReplaceAll(name, "\"", "")
	email = strings.ReplaceAll(email, "\"", "")

	switch strings.ToLower(action) {

	case "create":
		if email == "" {
			return errors.New("e-mail is required")
		}
		if _, err := mail.ParseAddress(email); err != nil {
			return errors.New("invalid E-mail format")
		}
		if !checkLetters.MatchString(name) {
			return errors.New("invalid name")
		}
	default:
		if email == "" {
			return errors.New("e-mail is required")
		}
		if _, err := mail.ParseAddress(email); err != nil {
			return errors.New("invalid E-mail format")
		}
		if !checkLetters.MatchString(name) {
			return errors.New("invalid name")
		}
	}
	return nil
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
	//var err error

	err := db.Debug().Model(&User{}).Where("email = ?", mail).Find(&u).Error
	return err

}

func ExtractToken(r *http.Request) string {
	tokenName, err := r.Cookie("token")

	if err != nil {
		return ""
	}
	return tokenName.Value
}

func ValidateToken(w http.ResponseWriter, r *http.Request) error {
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
