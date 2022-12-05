package controllers

import (
	"cadet-project/configurations"
	"cadet-project/models"
	"cadet-project/responses"
	"fmt"
	"github.com/crewjam/saml/samlsp"
	"net/http"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}
	}()

	userEmail := samlsp.AttributeFromContext(r.Context(), configurations.Config.Email)

	userName := samlsp.AttributeFromContext(r.Context(), configurations.Config.DisplayName)

	if userEmail == "" {
		return
	}
	user := models.User{
		Email: userEmail,
	}

	cookie := models.SetCookieToAllEndPoints(r)
	http.SetCookie(w, &cookie)

	userCreated, err := user.SaveUserDb(s.DB)
	if err != nil {
		responses.ERROR(w, 401, err)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))

	responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %v is authorized And created in database with Id: %v "+
		"And E-mail: %v, Token: %v, Display name: %v", userEmail, userCreated.ID, userCreated.Email, cookie.Value, userName))

}
