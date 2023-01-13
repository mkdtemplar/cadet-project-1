package controllers

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

func NewUserPrefController(IUserPreferencesRepository interfaces.IUserPreferencesRepository) *Server {
	return &Server{IUserPreferencesRepository: IUserPreferencesRepository}
}

func (s *Server) CreateUserPreferences(w http.ResponseWriter, r *http.Request) {
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

	err = validation.ValidateUserPref(userPref.UserCountry, userPref.UserId)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userPreferencesStore := validation.NewUserPrefObject(generate_id.GenerateID(), userPref.UserCountry, userPref.UserId)

	_, err = s.IUserPreferencesRepository.SaveUserPreferences(r.Context(), &userPreferencesStore)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, userPreferencesStore)
}

func (s *Server) GetUserPreference(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

	userPreferences, err := s.IUserPreferencesRepository.FindUserPreferences(r.Context(), id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userPreferences)
}

func (s *Server) GetUserPorts(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query().Get("id")
	paramsID, err := uuid.Parse(queryString)
	if err != nil {

		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in id format must be uuid"))
		return
	}

	userPreferences, err := s.IUserPreferencesRepository.FindUserPreferences(r.Context(), paramsID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	userPrefPorts, err := s.IUserPreferencesRepository.FindUserPrefPorts(r.Context(), userPreferences)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userPrefPorts)
}

func (s *Server) UpdateUserPreferences(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query().Get("id")
	paramsID, err := uuid.Parse(queryString)
	if err != nil {

		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in id format must be uuid"))
		return
	}

	userPrefFind, err := s.IUserPreferencesRepository.FindUserPreferences(r.Context(), paramsID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("user preferences not found"))
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var userPrefUpdate models.UserPreferences

	err = json.Unmarshal(body, &userPrefUpdate)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = validation.ValidateUserPref(userPrefUpdate.UserCountry, userPrefUpdate.UserId)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("data format validation failed"))
		return
	}

	userPrefUpdate = validation.NewUserPrefObject(userPrefFind.ID, userPrefUpdate.UserCountry, userPrefFind.UserId)

	_, err = s.IUserPreferencesRepository.UpdateUserPref(r.Context(), paramsID, userPrefUpdate.UserCountry)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, userPrefUpdate)
}

func (s *Server) DeleteUserPref(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query().Get("id")
	paramsID, err := uuid.Parse(queryString)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if _, err = s.IUserPreferencesRepository.DeleteUserPreferences(r.Context(), paramsID); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}
