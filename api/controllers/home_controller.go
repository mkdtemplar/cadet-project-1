package controllers

import (
	"cadet-project/configurations"
	"cadet-project/models"
	"cadet-project/responses"
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

	if userEmail == "" {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("user email not provided"))

	}
	user := models.User{
		Email: userEmail,
		Name:  userName,
	}
	user.PrepareUserData()
	err = user.ValidateUserData("")
	if err != nil {
		responses.ERROR(w, 401, errors.New("invalid E-mail format"))
	}
	tokenValue := models.ExtractToken(r)
	expiresAt := time.Now().Add(300 * time.Second)

	models.Sessions[tokenValue] = models.Session{Expiry: expiresAt}

	models.Cookie = models.CreateCookieToAllEndPoints(tokenValue, expiresAt)

	err = user.CheckUser(s.DB, userEmail)
	if err != nil {
		userCreated, err := user.SaveUserDb(s.DB)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))

		responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %v  with E-mail: %v is authorized and created in database", userName, userEmail))
	} else {
		responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %v  with E-mail: %v is already in database and authorized", userName, userEmail))
	}

}
