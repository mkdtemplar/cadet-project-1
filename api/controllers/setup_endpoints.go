package controllers

import (
	"cadet-project/configurations"
	"cadet-project/responses"
	"errors"
	"net/http"
)

func (s *Server) notFound(w http.ResponseWriter) {
	responses.ERROR(w, http.StatusInternalServerError, errors.New("path not found"))
	return
}

func (s *Server) ServeEndPoints(w http.ResponseWriter, r *http.Request) {
	configurations.InitConfig("configurations")

	userController, userPrefController := s.ControllersConstructor()

	w.Header().Set("content-type", "application/json")

	currentPath := r.URL.Path

	switch currentPath {
	case configurations.Config.UserDelete:
		userController.Delete(w, r, GetQueryID(w, r))
		return
	case configurations.Config.UserCreate:
		userController.Create(w, r)
		return
	case configurations.Config.UserPref:
		if r.Method == http.MethodPost {
			userPrefController.CreateUserPreferences(w, r)
		}
		if r.Method == http.MethodGet {
			userPrefController.GetUserPreference(w, r, GetQueryID(w, r))
		}
		if r.Method == http.MethodPatch {
			userPrefController.UpdateUserPreferences(w, r, GetQueryID(w, r))
		}
		if r.Method == http.MethodDelete {
			userPrefController.DeleteUserPref(w, r, GetQueryID(w, r))
		}
		return
	case configurations.Config.UserPorts:
		userPrefController.GetUserPorts(w, r, GetQueryID(w, r))
		return
	default:
		s.notFound(w)
		return
	}
}
