package controllers

import (
	"cadet-project/configurations"
	"cadet-project/controllers/helper"
	"cadet-project/interfaces"
	"cadet-project/middlewares_token_validation"
	"cadet-project/models"
	"cadet-project/repository"
	"cadet-project/repository/generate_id"
	"cadet-project/responses"
	"cadet-project/saml_handler"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func NewUserController(IUserRepository interfaces.IUserRepository, IUserPreferencesRepository interfaces.IUserPreferencesRepository,
	IShipPortsRepository interfaces.IShipPortsRepository) *Server {
	return &Server{IUserRepository: IUserRepository, IUserPreferencesRepository: IUserPreferencesRepository, IShipPortsRepository: IShipPortsRepository}
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var err error
	v := repository.Validation{}
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

	_, err = s.IUserRepository.Get(r.Context(), user)
	if err == nil {

		user, err = s.IShipPortsRepository.FindUserPorts(r.Context(), user)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}

		responses.JSON(w, http.StatusCreated, user)

		return
	}

	userNew, err := s.IUserRepository.Create(r.Context(), user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %s  with E-mail: %s is authorized and created in database", userNew.Email, userNew.Name))

}

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	v := repository.Validation{}
	user := helper.ParseUserRequestBody(w, r)

	checkCredentials := v.ValidateUserEmail(user.Email).ValidateUserName(user.Name)

	if checkCredentials.Err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, checkCredentials.Err)
		return
	}

	if _, err := s.IUserRepository.Create(r.Context(), user); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, user)

}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

	if _, err := s.IUserRepository.Delete(r.Context(), id); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}

func (s *Server) GetId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	user, err := s.IUserRepository.GetById(r.Context(), id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}
