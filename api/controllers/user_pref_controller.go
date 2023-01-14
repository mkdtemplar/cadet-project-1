package controllers

import (
	"cadet-project/interfaces"
	"cadet-project/models"
	"cadet-project/repository"
	"cadet-project/repository/generate_id"
	"cadet-project/responses"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func NewUserPrefController(IUserPreferencesRepository interfaces.IUserPreferencesRepository) *Server {
	return &Server{IUserPreferencesRepository: IUserPreferencesRepository}
}

func (s *Server) CreateUserPreferences(w http.ResponseWriter, r *http.Request) {
	userPref := ParseUserPrefRequestBody(w, r)

	err := repository.ValidateUserPref(userPref.UserCountry, userPref.UserId)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("user preferences validation failed"))
		return
	}

	userPreferencesStore := repository.NewUserPrefObject(generate_id.GenerateID(), userPref.UserCountry, userPref.UserId)

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

func (s *Server) GetUserPorts(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

	userPreferences, err := s.IUserPreferencesRepository.FindUserPreferences(r.Context(), id)
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

func (s *Server) UpdateUserPreferences(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var err error
	userPrefFind := &models.UserPreferences{}
	userPrefFind, err = s.IUserPreferencesRepository.FindUserPreferences(r.Context(), id)

	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("user preferences not found"))
	}

	userPrefUpdate := ParseUserPrefRequestBody(w, r)

	err = repository.ValidateCountry(userPrefUpdate.UserCountry)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("data format validation failed"))
		return
	}

	userPrefFind, err = s.IUserPreferencesRepository.UpdateUserPref(r.Context(), id, userPrefUpdate.UserCountry)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	userPrefFind, err = s.IUserPreferencesRepository.FindUserPreferences(r.Context(), id)
	responses.JSON(w, http.StatusOK, userPrefFind)
}

func (s *Server) DeleteUserPref(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

	if _, err := s.IUserPreferencesRepository.DeleteUserPreferences(r.Context(), id); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}
