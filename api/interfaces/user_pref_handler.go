package interfaces

import "net/http"

type IUserPrefHandlers interface {
	CreateUserPreferences(w http.ResponseWriter, r *http.Request)
	GetUserPreference(w http.ResponseWriter, r *http.Request)
	GetAllUserPreferences(w http.ResponseWriter, r *http.Request)
	GetUserPorts(w http.ResponseWriter, r *http.Request)
	GetAllUserPorts(w http.ResponseWriter, r *http.Request)
	UpdateUserPreferences(w http.ResponseWriter, r *http.Request)
	DeleteUserPref(w http.ResponseWriter, r *http.Request)
}
