package handlers

import (
	"cadet-project/interfaces"
	"cadet-project/models"
	"cadet-project/repository/generate_id"
	"cadet-project/responses"
	"cadet-project/validation"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type UserPrefHandler struct {
	userPreferences interfaces.IUserPreferences
}

func NewUserPrefHandler(usr interfaces.IUserPreferences) interfaces.IUserPrefHandlers {
	return &UserPrefHandler{userPreferences: usr}
}

func (u *UserPrefHandler) CreateUserPreferences(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		userPref := models.UserPreferences{}

		err = json.Unmarshal(body, &userPref)

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		err = validation.ValidateUserPref("create", userPref.Country, userPref.UserId)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		userPrefCreated := models.UserPreferences{
			ID:      generate_id.GenerateID(),
			Country: userPref.Country,
			UserId:  userPref.UserId,
		}
		validation.ConstructUserPrefObject(userPref.Country)

		err = u.userPreferences.SaveUserPreferences(r.Context(), &userPrefCreated)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusCreated, userPrefCreated)
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
	}
}

func (u *UserPrefHandler) GetUserPreference(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryString := r.URL.Query().Get("user_id")
		paramsID, err := uuid.Parse(queryString)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in user_id format must be integer"))
			return
		}

		userPreferences, err := u.userPreferences.FindUserPreferences(r.Context(), paramsID)

		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, userPreferences)
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
	}
}

func (u *UserPrefHandler) GetUserPorts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryString := r.URL.Query().Get("user_id")
		paramsID, err := uuid.Parse(queryString)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in user_id format must be integer"))
			return
		}

		userPreferences, err := u.userPreferences.FindUserPreferences(r.Context(), paramsID)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		userPrefPorts, err := u.userPreferences.FindUserPrefPorts(r.Context(), userPreferences)

		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, userPrefPorts)
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
	}
}

func (u *UserPrefHandler) UpdateUserPreferences(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		queryString := r.URL.Query().Get("user_id")
		paramsID, err := uuid.Parse(queryString)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in user_id format must be integer"))
			return
		}

		userPrefFind, err := u.userPreferences.FindUserPreferences(r.Context(), paramsID)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, errors.New("user preferences not found"))
			return
		}

		body, err := io.ReadAll(r.Body)

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		userPrefUpdate := models.UserPreferences{}

		err = json.Unmarshal(body, &userPrefUpdate)

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		err = validation.ValidateUserPref("update", userPrefUpdate.Country, userPrefUpdate.UserId)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("data format validation failed"))
			return
		}

		userPrefUpdate.UserId = userPrefFind.UserId
		userPrefUpdate.ID = userPrefFind.ID
		validation.ConstructUserPrefObject(userPrefUpdate.Country)

		_, err = u.userPreferences.UpdateUserPref(r.Context(), paramsID, userPrefUpdate.Country)

		if err != nil {

			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, userPrefUpdate)
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
	}
}

func (u *UserPrefHandler) DeleteUserPref(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		queryString := r.URL.Query().Get("user_id")
		paramsID, err := uuid.Parse(queryString)
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		if _, err = u.userPreferences.DeleteUserPreferences(r.Context(), paramsID); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusNoContent, "")
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid http method"))
		return
	}
}
