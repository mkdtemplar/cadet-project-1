package interfaces

import (
	"net/http"

	"github.com/google/uuid"
)

type IUserPrefController interface {
	CreateUserPreferences(w http.ResponseWriter, r *http.Request)
	GetUserPreference(w http.ResponseWriter, r *http.Request, id uuid.UUID)
	GetAllUserPreferences(w http.ResponseWriter, r *http.Request, id uuid.UUID)
	GetUserPorts(w http.ResponseWriter, r *http.Request)
	UpdateUserPreferences(w http.ResponseWriter, r *http.Request)
	DeleteUserPref(w http.ResponseWriter, r *http.Request)
}
