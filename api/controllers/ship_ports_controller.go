package controllers

import (
	"cadet-project/interfaces"
	"cadet-project/responses"
	"net/http"

	"github.com/google/uuid"
)

func NewShipPortsController(IShipPortsRepository interfaces.IShipPortsRepository) *Server {
	return &Server{IShipPortsRepository: IShipPortsRepository}
}

func (s *Server) GetUserPorts(w http.ResponseWriter, r *http.Request, id uuid.UUID) {

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
