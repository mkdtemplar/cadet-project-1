package interfaces

import (
	"net/http"

	"github.com/google/uuid"
)

type IUserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request, id uuid.UUID)
	GetId(w http.ResponseWriter, r *http.Request, id uuid.UUID)
}
