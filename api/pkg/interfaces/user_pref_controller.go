package interfaces

import (
	"net/http"

	"github.com/google/uuid"
)

type IUserPrefController interface {
	CreateUserPreferences(w http.ResponseWriter, r *http.Request)
	GetUserPreference(w http.ResponseWriter, r *http.Request, id uuid.UUID)

	UpdateUserPreferences(w http.ResponseWriter, r *http.Request, userid uuid.UUID)
	DeleteUserPref(w http.ResponseWriter, r *http.Request, id uuid.UUID)
}
