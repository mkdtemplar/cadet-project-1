package interfaces

import (
	"net/http"

	"github.com/google/uuid"
)

type IShipController interface {
	GetUserPorts(w http.ResponseWriter, r *http.Request, id uuid.UUID)
	GetPorts(w http.ResponseWriter, r *http.Request, id uuid.UUID)
}
