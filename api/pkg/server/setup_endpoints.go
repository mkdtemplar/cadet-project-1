package server

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers/helper"
	"cadet-project/pkg/responses"
	"errors"
	"net/http"
)

func (s *Server) notFound(w http.ResponseWriter) {
	responses.ERROR(w, http.StatusInternalServerError, errors.New("path not found"))
	return
}

func (s *Server) ServeUserPrefEndPoints(w http.ResponseWriter, r *http.Request) {

	userPrefController := C.Controllers()
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
	config.InitConfig("configurations")

	controllers := C.Controllers()

	w.Header().Set("content-type", "application/json")

	currentPath := r.URL.Path

	switch currentPath {
	case config.Config.UserDelete:
		controllers.Delete(w, r, helper.GetQueryID(w, r))
		return
	case config.Config.UserCreate:
		controllers.Create(w, r)
		return
	case config.Config.UserPref:
		s.ServeUserPrefEndPoints(w, r)
		return
	case config.Config.UserPorts:
		controllers.GetUserPorts(w, r, helper.GetQueryID(w, r))
		return
	case config.Config.UserPrefPorts:
		controllers.GetUserPrefPorts(w, r, helper.GetQueryID(w, r))
		return
	case config.Config.UserId:
		controllers.GetId(w, r, helper.GetQueryID(w, r))
		return
	default:
		s.notFound(w)
		return
	}
}
