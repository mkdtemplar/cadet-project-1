package controllers

import (
	"cadet-project/responses"
	"errors"
	"net/http"
)

func (s *Server) notFound(w http.ResponseWriter) {
	responses.ERROR(w, http.StatusInternalServerError, errors.New("path not found"))
	return
}
func (s *Server) ServeEndPoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodDelete && r.URL.Path == "/userdelete":
		s.DeleteUser(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/usercreate":
		s.CreateUserInDb(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/userpref":
		s.CreateUserPreferences(w, r)
		return
	case r.Method == http.MethodGet && r.URL.Path == "/userpref":
		s.GetUserPreference(w, r)
		return
	case r.Method == http.MethodGet && r.URL.Path == "/userprefports":
		s.GetUserPorts(w, r)
		return
	case r.Method == http.MethodDelete && r.URL.Path == "/userpref":
		s.DeleteUserPref(w, r)
		return
	case r.Method == http.MethodPut && r.URL.Path == "/userpref":
		s.UpdateUserPreferences(w, r)
	default:
		s.notFound(w)
		return
	}
}
