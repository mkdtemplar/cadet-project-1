package handlers

import (
	"cadet-project/repository"
	"cadet-project/responses"
	"errors"
	"net/http"
)

func (s *Server) notFound(w http.ResponseWriter) {
	responses.ERROR(w, http.StatusInternalServerError, errors.New("path not found"))
	return
}
func (s *Server) ServeEndPoints(w http.ResponseWriter, r *http.Request) {
	usrRepo := repository.NewUserRepo(s.DB)
	usrHandlers := NewUserHandler(usrRepo)
	userPrefRepo := repository.NewUserPrefRepo(s.DB)
	userPrefHandlers := NewUserPrefHandler(userPrefRepo)

	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodDelete && r.URL.Path == "/userdelete":
		usrHandlers.DeleteUser(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/usercreate":
		usrHandlers.CreateUserInDb(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/userpref":
		userPrefHandlers.CreateUserPreferences(w, r)
	case r.Method == http.MethodGet && r.URL.Path == "/userpref":
		userPrefHandlers.GetUserPreference(w, r)
		return
	case r.Method == http.MethodGet && r.URL.Path == "/userprefports":
		userPrefHandlers.GetUserPorts(w, r)
		return
	case r.Method == http.MethodDelete && r.URL.Path == "/userpref":
		userPrefHandlers.DeleteUserPref(w, r)
		return
	case r.Method == http.MethodPatch && r.URL.Path == "/userpref":
		userPrefHandlers.UpdateUserPreferences(w, r)
	default:
		s.notFound(w)
		return
	}
}
