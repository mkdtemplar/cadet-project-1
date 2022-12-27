package controllers

import (
	"cadet-project/configurations"
	"cadet-project/models"
	"cadet-project/responses"
	"cadet-project/validation"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/crewjam/saml/samlsp"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	var err error

	userEmail := samlsp.AttributeFromContext(r.Context(), configurations.Config.Email)

	userName := samlsp.AttributeFromContext(r.Context(), configurations.Config.DisplayName)
	err = validation.ValidateUserData("create", userEmail, userName)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("invalid user email format"))
		return
	}
	user := models.User{
		Email: userEmail,
		Name:  userName,
	}

	tokenValue := validation.ExtractToken(r)
	expiresAt := time.Now().Add(300 * time.Second)

	models.Sessions[tokenValue] = models.Session{Expiry: expiresAt}

	models.Cookie.Expires = expiresAt
	models.Cookie.Path = "/"
	http.SetCookie(w, &models.Cookie)

	err = user.CheckUser(s.DB, userEmail)
	if err != nil {
		_, err = user.SaveUserDb(s.DB)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %s  with E-mail: %s is authorized and created in database", userName, userEmail))
	} else {
		responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %s  with E-mail: %s is already in database and authorized", userName, userEmail))
	}

}
