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
	var UserHandlers = s.UserHandlerFunc()

	var UserPrefHandlers = s.UserPrefHandlerFunc()

	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodDelete && r.URL.Path == "/userdelete":
		UserHandlers.DeleteUser(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/usercreate":
		UserHandlers.CreateUserInDb(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/userpref":
		UserPrefHandlers.CreateUserPreferences(w, r)
	case r.Method == http.MethodGet && r.URL.Path == "/userpref":
		UserPrefHandlers.GetUserPreference(w, r)
		return
	case r.Method == http.MethodGet && r.URL.Path == "/userprefports":
		UserPrefHandlers.GetUserPorts(w, r)
		return
	case r.Method == http.MethodDelete && r.URL.Path == "/userpref":
		UserPrefHandlers.DeleteUserPref(w, r)
		return
	case r.Method == http.MethodPatch && r.URL.Path == "/userpref":
		UserPrefHandlers.UpdateUserPreferences(w, r)
	default:
		s.notFound(w)
		return
	}
}
