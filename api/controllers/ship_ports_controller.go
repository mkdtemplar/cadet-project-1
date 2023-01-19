package controllers

import (
	"cadet-project/interfaces"
	"cadet-project/responses"
	"net/http"

	"github.com/google/uuid"
)

func NewShipPortsController(IUserRepository interfaces.IUserRepository, IUserPreferencesRepository interfaces.IUserPreferencesRepository, IShipPortsRepository interfaces.IShipPortsRepository) *Server {
	return &Server{IUserRepository: IUserRepository, IUserPreferencesRepository: IUserPreferencesRepository, IShipPortsRepository: IShipPortsRepository}
}

func (s *Server) GetUserPrefPorts(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

	userPreferences, err := s.IUserPreferencesRepository.FindUserPreferences(r.Context(), id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	userPrefPorts, err := s.IShipPortsRepository.FindUserPrefPorts(r.Context(), userPreferences)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userPrefPorts)
}

func (s *Server) GetUserPorts(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	user, err := s.IUserRepository.GetById(r.Context(), id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	userPorts, err := s.IShipPortsRepository.FindUserPorts(r.Context(), user)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userPorts)
}
