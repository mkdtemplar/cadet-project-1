package controllers

import (
	"cadet-project/configurations"
	"cadet-project/controllers/helper"
	"cadet-project/responses"
	"errors"
	"net/http"
)

func (s *Server) notFound(w http.ResponseWriter) {
	responses.ERROR(w, http.StatusInternalServerError, errors.New("path not found"))
	return
}

func (s *Server) ServeUserPrefEndPoints(w http.ResponseWriter, r *http.Request) {

	_, userPrefController, _ := s.ControllersConstructor()
	switch {
	case r.Method == http.MethodPost:
		userPrefController.CreateUserPreferences(w, r)
		return
	case r.Method == http.MethodGet:
		userPrefController.GetUserPreference(w, r, helper.GetQueryID(w, r))
		return
	case r.Method == http.MethodPatch:
		userPrefController.UpdateUserPreferences(w, r, helper.GetQueryID(w, r))
		return
	case r.Method == http.MethodDelete:
		userPrefController.DeleteUserPref(w, r, helper.GetQueryID(w, r))
		return
	default:
		s.notFound(w)
		return
	}
}

func (s *Server) ServeEndPoints(w http.ResponseWriter, r *http.Request) {
	configurations.InitConfig("configurations")

	userController, _, shipPortsController := s.ControllersConstructor()

	w.Header().Set("content-type", "application/json")

	currentPath := r.URL.Path

	switch currentPath {
	case configurations.Config.UserDelete:
		userController.Delete(w, r, helper.GetQueryID(w, r))
		return
	case configurations.Config.UserCreate:
		userController.Create(w, r)
		return
	case configurations.Config.UserPref:
		s.ServeUserPrefEndPoints(w, r)
		return
	case configurations.Config.UserPorts:
		shipPortsController.GetUserPorts(w, r, helper.GetQueryID(w, r))
		return
	case configurations.Config.UserPrefPorts:
		shipPortsController.GetUserPrefPorts(w, r, helper.GetQueryID(w, r))
	case configurations.Config.UserId:
		userController.GetId(w, r, helper.GetQueryID(w, r))
	default:
		s.notFound(w)
		return
	}
}
