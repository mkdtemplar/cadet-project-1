package controllers

import (
	"cadet-project/configurations"
	"cadet-project/interfaces"
	"cadet-project/middlewares_token_validation"
	"cadet-project/models"
	"cadet-project/repository/generate_id"
	"cadet-project/repository/validation"
	"cadet-project/responses"
	"cadet-project/saml_handler"
	"fmt"
	"net/http"
	"time"
)

func NewLoginController(IUserRepository interfaces.IUserRepository, IUserPreferencesRepository interfaces.IUserPreferencesRepository, IShipPortsRepository interfaces.IShipPortsRepository) *Server {
	return &Server{IUserRepository: IUserRepository, IUserPreferencesRepository: IUserPreferencesRepository, IShipPortsRepository: IShipPortsRepository}
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var err error
	v := validation.Validation{}
	userEmail, userName := saml_handler.Credentials(w, r, configurations.Config.Email, configurations.Config.DisplayName)

	checkCredentials := v.ValidateUserEmail(userEmail).ValidateUserName(userName)

	if checkCredentials.Err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, checkCredentials.Err)
		return
	}

	user := &models.User{
		ID:    generate_id.GenerateID(),
		Email: userEmail,
		Name:  userName,
	}

	tokenValue := middlewares_token_validation.ExtractToken(r)
	expiresAt := time.Now().Add(900 * time.Second)

	models.Sessions[tokenValue] = models.Session{Expiry: expiresAt}

	models.Cookie.Expires = expiresAt
	models.Cookie.Path = "/"
	http.SetCookie(w, &models.Cookie)

	checkUser, err := s.IUserRepository.GetUserEmail(r.Context(), userEmail)
	if err == nil {
		userPorts, err := s.IShipPortsRepository.FindUserPorts(r.Context(), checkUser.ID)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
		responses.JSON(w, http.StatusCreated, userPorts)
		return
	}

	userNew, err := s.IUserRepository.Create(r.Context(), user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %s  with E-mail: %s is authorized and created in database", userNew.Email, userNew.Name))

}