package interfaces

import (
	"net/http"

	"github.com/google/uuid"
)

type IUserController interface {
	Home(w http.ResponseWriter, r *http.Request)
	CreateUserInDb(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request, id uuid.UUID)
}
