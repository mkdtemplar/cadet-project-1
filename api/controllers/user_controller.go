package controllers

import (
	"cadet-project/configurations"
	"cadet-project/interfaces"
	"cadet-project/saml_handler"

	"cadet-project/models"
	"cadet-project/repository/generate_id"
	"cadet-project/responses"
	"cadet-project/validation"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func NewUserController(IUserRepository interfaces.IUserRepository) *Server {
	return &Server{IUserRepository: IUserRepository}
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	var err error

	userEmail, userName := saml_handler.Credentials(w, r, configurations.Config.Email, configurations.Config.DisplayName)

	user := &models.User{
		ID:    generate_id.GenerateID(),
		Email: userEmail,
		Name:  userName,
	}

	tokenValue := validation.ExtractToken(r)
	expiresAt := time.Now().Add(900 * time.Second)

	models.Sessions[tokenValue] = models.Session{Expiry: expiresAt}

	models.Cookie.Expires = expiresAt
	models.Cookie.Path = "/"
	http.SetCookie(w, &models.Cookie)

	_, err = s.IUserRepository.GetUser(r.Context(), user)
	if err == nil {
		responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %s  with E-mail: %s is already in database and authorized", userName, userEmail))
		return
	}

	userNew, err := s.IUserRepository.SaveUserDb(r.Context(), user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %s  with E-mail: %s is authorized and created in database", userNew.Email, userNew.Name))

}

func (s *Server) TestCreate() {
	s.TestCreateUser()
}

func (s *Server) TestCreate1() {
	s.Name(s.IUserRepository.SaveUserDb)
}

func (s *Server) CreateUserInDb(w http.ResponseWriter, r *http.Request) {

	user := RequestBodyUser(w, r)

	err := validation.ValidateUserData(user.Email, user.Name)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("invalid user email format"))
		return
	}
	s.IUserRepository.PrepareUserData(user.Email, user.Name)
	if _, err = s.IUserRepository.SaveUserDb(r.Context(), &user); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, user)

}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

	if _, err := s.IUserRepository.DeleteUserDb(r.Context(), id); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}
