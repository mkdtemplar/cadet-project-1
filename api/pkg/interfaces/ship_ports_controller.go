package interfaces

import (
	"net/http"

	"github.com/google/uuid"
)

type IShipController interface {
	GetUserPrefPorts(w http.ResponseWriter, r *http.Request, id uuid.UUID)
	GetUserPorts(w http.ResponseWriter, r *http.Request, id uuid.UUID)
}
