package handlers

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

	userConstructor, userPrefConstructor := s.HandlersConstructor()

	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodDelete && r.URL.Path == "/userdelete":
		userConstructor.DeleteUser(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/usercreate":
		userConstructor.CreateUserInDb(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/userpref":
		userPrefConstructor.CreateUserPreferences(w, r)
	case r.Method == http.MethodGet && r.URL.Path == "/userpref":
		userPrefConstructor.GetUserPreference(w, r)
		return
	case r.Method == http.MethodGet && r.URL.Path == "/userprefports":
		userPrefConstructor.GetUserPorts(w, r)
		return
	case r.Method == http.MethodDelete && r.URL.Path == "/userpref":
		userPrefConstructor.DeleteUserPref(w, r)
		return
	case r.Method == http.MethodPatch && r.URL.Path == "/userpref":
		userPrefConstructor.UpdateUserPreferences(w, r)
	default:
		s.notFound(w)
		return
	}
}
