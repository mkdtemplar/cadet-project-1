package handlers

import (
	"cadet-project/configurations"
	"cadet-project/models"
	"cadet-project/repository/interfaces"
	"cadet-project/responses"
	"cadet-project/validation"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/crewjam/saml/samlsp"
)

type IUserHandlers interface {
	CreateUserInDb(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	Home(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	user interfaces.IUserRepository
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
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

	_, err = h.user.CheckUser(r.Context(), &user)
	if err != nil {
		err = h.user.SaveUserDb(r.Context(), &user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %s  with E-mail: %s is authorized and created in database", userName, userEmail))
	} else {
		responses.JSON(w, http.StatusCreated, fmt.Sprintf("User : %s  with E-mail: %s is already in database and authorized", userName, userEmail))
	}
}

func (h *Handler) CreateUserInDb(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}

		user := models.User{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		err = validation.ValidateUserData("create", user.Email, user.Name)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("invalid user email format"))
			return
		}
		h.user.PrepareUserData(user.Email, user.Name)
		if err = h.user.SaveUserDb(r.Context(), &user); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusCreated, user)
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
		return
	}
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		paramsID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		if _, err = h.user.DeleteUserDb(r.Context(), paramsID); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusNoContent, "")
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
		return
	}
}

func NewUserHandler(usr interfaces.IUserRepository) IUserHandlers {
	return &Handler{user: usr}
}
