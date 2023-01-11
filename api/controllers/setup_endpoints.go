package controllers

import (
	"cadet-project/configurations"
	"cadet-project/responses"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func getQueryID(w http.ResponseWriter, r *http.Request) uuid.UUID {
	queryString := r.URL.Query().Get("id")
	paramsID, err := uuid.Parse(queryString)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in id format must be uuid"))
		return uuid.Nil
	}
	return paramsID
}

func getQueryUserId(w http.ResponseWriter, r *http.Request) uuid.UUID {
	queryString := r.URL.Query().Get("user_id")
	paramsID, err := uuid.Parse(queryString)
	if err != nil {

		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error in id format must be uuid"))
		return uuid.Nil
	}
	return paramsID
}

func (s *Server) notFound(w http.ResponseWriter) {
	responses.ERROR(w, http.StatusInternalServerError, errors.New("path not found"))
	return
}

func (s *Server) ServeEndPoints(w http.ResponseWriter, r *http.Request) {
	configurations.InitConfig("configurations")

	userController, userPrefController := s.ControllerConstructor()

	w.Header().Set("content-type", "application/json")

	currentPath := r.URL.Path

	switch currentPath {
	case configurations.Config.UserDelete:
		userController.DeleteUser(w, r)
		return
	case configurations.Config.UserCreate:
		userController.CreateUserInDb(w, r)
		return
	case configurations.Config.UserPref:
		if r.Method == http.MethodPost {
			userPrefController.CreateUserPreferences(w, r)
		}
		if r.Method == http.MethodGet {
			userPrefController.GetUserPreference(w, r, getQueryID(w, r))
		}
		if r.Method == http.MethodPatch {
			userPrefController.UpdateUserPreferences(w, r)
		}
		if r.Method == http.MethodDelete {
			userPrefController.DeleteUserPref(w, r)
		}
		return
	case configurations.Config.UserPorts:
		userPrefController.GetUserPorts(w, r)
		return
	case configurations.Config.ListUserPref:
		userPrefController.GetAllUserPreferences(w, r, getQueryUserId(w, r))
	default:
		s.notFound(w)
		return
	}
}
