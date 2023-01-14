package repository

import (
	"cadet-project/interfaces"
	"cadet-project/models"
	"context"
	"errors"
	"html"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ValidateCountry(country string) error {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]+$`)
	if checkLetters.MatchString(country) == false {
		return errors.New("country string wrong format")
	}

	return nil
}
func ValidateUserPref(country string, userId uuid.UUID) error {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]+$`)
	checkId := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

	if checkLetters.MatchString(country) == false {
		return errors.New("country string wrong format")
	}
	if userId == uuid.Nil || checkId.MatchString(userId.String()) == false {
		return errors.New("user id is required or wrong data format user_id must be uuid")
	}

	return nil
}

func NewUserPrefObject(id uuid.UUID, country string, userId uuid.UUID) models.UserPreferences {
	userPref := models.UserPreferences{}
	country = html.EscapeString(strings.TrimSpace(country))

	userPref = models.UserPreferences{
		ID:          id,
		UserCountry: country,
		UserId:      userId,
	}
	return userPref
}

func NewUserPrefRepo(db *gorm.DB) interfaces.IUserPreferencesRepository {
	return &PG{DB: db}
}

func (u *PG) GetAllUserPreferences(ctx context.Context, userId uuid.UUID) ([]*models.UserPreferences, error) {
	var userPrefList []*models.UserPreferences

	err := u.DB.WithContext(ctx).Model(&models.UserPreferences{}).Where("user_id = ?", userId).Find(&userPrefList).Error
	if err != nil {
		return nil, err
	}

	return userPrefList, nil
}

func (u *PG) SaveUserPreferences(ctx context.Context, in *models.UserPreferences) (*models.UserPreferences, error) {
	if in == nil {
		return &models.UserPreferences{}, errors.New("user details empty")
	}

	err := u.DB.WithContext(ctx).Model(u.UserPreferences).Create(&in).Error
	if err != nil {
		return &models.UserPreferences{}, errors.New("user not created")
	}
	return in, nil
}

func (u *PG) FindUserPreferences(ctx context.Context, id uuid.UUID) (*models.UserPreferences, error) {
	var err error

	userPref := &models.UserPreferences{}

	err = u.DB.WithContext(ctx).Model(&models.UserPreferences{}).Where("id = ?", id).Find(&userPref).Error
	if err != nil {
		return nil, err
	}

	return userPref, nil
}

func (u *PG) UpdateUserPref(ctx context.Context, userid uuid.UUID, country string) (*models.UserPreferences, error) {
	var err error

	err = u.DB.WithContext(ctx).Model(&models.UserPreferences{}).
		Where("id = ?", userid).Take(&models.UserPreferences{}).UpdateColumns(map[string]interface{}{"user_country": country}).Error
	if err != nil {
		err.Error()
		return nil, err
	}

	return u.UserPreferences, nil

}

func (u *PG) DeleteUserPreferences(ctx context.Context, id uuid.UUID) (int64, error) {
	var err error

	tx := u.DB.Begin()

	delTx := tx.WithContext(ctx).Model(&models.UserPreferences{}).Where("id = ?", id).Delete(&models.UserPreferences{})

	if err = delTx.Error; err != nil {
		return 0, err
	} else {
		tx.Commit()
	}

	return delTx.RowsAffected, nil
}

func (u *PG) FindUserPrefPorts(ctx context.Context, in *models.UserPreferences) (*models.UserPreferences, error) {
	var err error
	userPref := &models.UserPreferences{}
	if err = u.DB.WithContext(ctx).Joins("join ship_ports ON ship_ports.country = user_preferences.user_country").Find(userPref).Error; err != nil {
		return nil, err
	}

	var ports []models.ShipPorts

	if err := u.DB.WithContext(ctx).Where("country = ?", in.UserCountry).Model(&models.ShipPorts{}).Find(&ports).Error; err != nil {
		return &models.UserPreferences{}, err
	}

	userPref.Ports = ports

	return userPref, nil
}
