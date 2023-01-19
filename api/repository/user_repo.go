package repository

import (
	"cadet-project/interfaces"
	"cadet-project/models"
	"cadet-project/repository/generate_id"
	"context"
	"errors"
	"html"
	"net/mail"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PG struct {
	DB              *gorm.DB
	User            *models.User
	UserPreferences *models.UserPreferences
	ShipPorts       *models.ShipPorts
}

func (u *PG) FindUserPorts(ctx context.Context, usr *models.User) (*models.User, error) {
	var err error

	user := models.User{}
	if err = u.DB.WithContext(ctx).Model(&models.User{}).Joins("join user_preferences ON user_preferences.user_id = users.id").Find(&user).Error; err != nil {
		return nil, err
	}
	var userPref models.UserPreferences
	if err = u.DB.WithContext(ctx).Where("user_id = ?", usr.ID).Model(&models.UserPreferences{}).Find(&userPref).Error; err != nil {
		return nil, err
	}

	if err = u.DB.WithContext(ctx).Joins("join ship_ports ON ship_ports.country = user_preferences.user_country").Find(&userPref).Error; err != nil {
		return nil, err
	}

	var ports []models.ShipPorts

	if err := u.DB.WithContext(ctx).Where("country = ?", userPref.UserCountry).Model(&models.ShipPorts{}).Find(&ports).Error; err != nil {
		return nil, err
	}

	userPref.Ports = ports
	user.UserPref = userPref
	return &user, nil
}

type Validation struct {
	Err error
}

func (v *Validation) Error() string {
	var err string
	return err
}

func (v *Validation) ValidateUserEmail(email string) *Validation {
	email = strings.ReplaceAll(email, "\"", "")
	email = strings.ToLower(email)
	email = html.EscapeString(strings.TrimSpace(email))

	if _, err := mail.ParseAddress(email); err != nil {
		v.Err = errors.New("invalid E-mail format")
		return v
	}
	return v
}

func (v *Validation) ValidateUserName(name string) *Validation {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]*$`)
	if !checkLetters.MatchString(name) {
		v.Err = errors.New("invalid name")
		return v
	}
	return v
}

func NewUserRepo(db *gorm.DB) interfaces.IUserRepository {
	return &PG{DB: db}
}

func (u *PG) PrepareUserData(email string, name string) {
	email = html.EscapeString(strings.TrimSpace(email))
	name = html.EscapeString(strings.TrimSpace(name))
}

func (u *PG) Create(ctx context.Context, usr *models.User) (*models.User, error) {
	if usr == nil {
		return &models.User{}, errors.New("user details empty")
	}
	usr.ID = generate_id.GenerateID()
	err := u.DB.WithContext(ctx).Model(u.User).Create(&usr).Error

	if err != nil {
		return &models.User{}, err
	}

	return usr, nil
}

func (u *PG) Delete(ctx context.Context, uid uuid.UUID) (int64, error) {
	var err error

	tx := u.DB.Begin()

	delTx := tx.WithContext(ctx).Model(&models.User{}).Delete(&models.User{}, uid)

	if err = delTx.Error; err != nil {
		return 0, err
	} else {
		tx.Commit()
	}

	return tx.RowsAffected, nil
}

func (u *PG) Get(ctx context.Context, in *models.User) (*models.User, error) {
	user := &models.User{}

	err := u.DB.WithContext(ctx).Model(&models.User{}).Take(user, "email = ?", in.Email).Error
	return user, err
}

func (u *PG) GetById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var err error

	user := &models.User{}

	err = u.DB.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
