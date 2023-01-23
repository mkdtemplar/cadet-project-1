package controllers

import (
	"cadet-project/pkg/config"
	interfaces2 "cadet-project/pkg/interfaces"
	"cadet-project/pkg/middlewares_token_validation"
	models2 "cadet-project/pkg/models"
	"cadet-project/pkg/repository/generate_id"
	"cadet-project/pkg/repository/validation"
	"cadet-project/pkg/responses"
	"cadet-project/pkg/saml_handler"
	"fmt"
	"net/http"
	"time"
)

func NewLoginController(IUserRepository interfaces2.IUserRepository, IUserPreferencesRepository interfaces2.IUserPreferencesRepository, IShipPortsRepository interfaces2.IShipPortsRepository) *Controller {
	return &Controller{IUserRepository: IUserRepository, IUserPreferencesRepository: IUserPreferencesRepository, IShipPortsRepository: IShipPortsRepository}
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var err error
	v := validation.Validation{}
	userEmail, userName := saml_handler.Credentials(w, r, config.Config.Email, config.Config.DisplayName)

	checkCredentials := v.ValidateUserEmail(userEmail).ValidateUserName(userName)

	if checkCredentials.Err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, checkCredentials.Err)
		return
	}

	user := &models2.User{
		ID:    generate_id.GenerateID(),
		Email: userEmail,
		Name:  userName,
	}

	tokenValue := middlewares_token_validation.ExtractToken(r)
	expiresAt := time.Now().Add(900 * time.Second)

	models2.Sessions[tokenValue] = models2.Session{Expiry: expiresAt}

	models2.Cookie.Expires = expiresAt
	models2.Cookie.Path = "/"
	http.SetCookie(w, &models2.Cookie)

	checkUser, err := c.IUserRepository.GetUserEmail(r.Context(), userEmail)
	if err == nil {
		userPorts, err := c.IShipPortsRepository.FindUserPorts(r.Context(), checkUser.ID)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
		responses.JSON(w, http.StatusCreated, userPorts)
		return
	}

	userNew, err := c.IUserRepository.Create(r.Context(), user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %s  with E-mail: %s is authorized and created in database", userNew.Email, userNew.Name))

}
