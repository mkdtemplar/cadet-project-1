package controllers

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/interfaces"
	"cadet-project/pkg/middlewares"
	"cadet-project/pkg/models"
	"cadet-project/pkg/repository/generate_id"
	"cadet-project/pkg/repository/validation"
	"cadet-project/pkg/responses"
	"cadet-project/pkg/saml_handler"
	"errors"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func NewLoginController(IUserRepository interfaces.IUserRepository, IShipPortsRepository interfaces.IShipPortsRepository) *LoginController {
	return &LoginController{IUserRepository: IUserRepository, IShipPortsRepository: IShipPortsRepository}
}

func (l *LoginController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var err error
	config.InitDbConfig("pkg/config")
	defer func() {
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}
	}()
	v := validation.Validation{}
	userEmail, userName := saml_handler.Credentials(w, r, config.Config.Email, config.Config.DisplayName)

	checkCredentials := v.ValidateUserEmail(userEmail).ValidateUserName(userName)
	err = checkCredentials.Err
	if err != nil {
		return
	}

	user := &models.User{
		ID:    generate_id.GenerateID(),
		Email: userEmail,
		Name:  userName,
	}

	tokenValue := middlewares.ExtractToken(r)
	expiresAt := time.Now().Add(900 * time.Second)

	models.AddSession(tokenValue, models.Session{Expiry: expiresAt})

	models.Cookie.Expires = expiresAt
	models.Cookie.Path = "/"
	http.SetCookie(w, &models.Cookie)
	var checkUser *models.User
	checkUser, err = l.IUserRepository.GetUserEmail(r.Context(), userEmail)
	if err == nil && checkUser != nil {
		var userPorts *models.User
		userPorts, err = l.IShipPortsRepository.FindUserPorts(r.Context(), checkUser.ID)
		if err != nil {
			return
		}
		responses.JSON(w, http.StatusOK, userPorts)
		return
	}

	var userNew *models.User
	if errors.Is(err, gorm.ErrRecordNotFound) {
		userNew, err = l.IUserRepository.Create(r.Context(), user)
		if err != nil {
			return
		}
	} else {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("problem with database or http server"))
	}

	responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %s  with E-mail: %s is authorized and created in database", userNew.Email, userNew.Name))
}
