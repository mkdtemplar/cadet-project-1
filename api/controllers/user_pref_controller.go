package controllers

import (
	"cadet-project/interfaces"
	"cadet-project/repository/generate_id"
	"cadet-project/responses"
	"cadet-project/validation"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func NewUserPrefController(IUserPreferencesRepository interfaces.IUserPreferencesRepository) *Server {
	return &Server{IUserPreferencesRepository: IUserPreferencesRepository}
}

func (s *Server) CreateUserPreferences(w http.ResponseWriter, r *http.Request) {
	userPref := RequestBodyUserPref(w, r)

	err := validation.ValidateUserPref("create", userPref.UserCountry, userPref.UserId)
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

func (s *Server) UpdateUserPreferences(w http.ResponseWriter, r *http.Request, userid uuid.UUID) {

	userPrefFind, err := s.IUserPreferencesRepository.FindUserPreferences(r.Context(), userid)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("user preferences not found"))
		return
	}

	userPrefUpdate := RequestBodyUserPref(w, r)

	err = validation.ValidateCountry(userPrefUpdate.UserCountry)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("data format validation failed"))
		return
	}

	userPrefUpdate.ID = userPrefFind.ID
	userPrefUpdate.UserId = userPrefFind.UserId

	_, err = s.IUserPreferencesRepository.UpdateUserPref(r.Context(), userid, userPrefUpdate.UserCountry)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, userPrefUpdate)
}

func (s *Server) DeleteUserPref(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

	if _, err := s.IUserPreferencesRepository.DeleteUserPreferences(r.Context(), id); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}
