package interfaces

import "net/http"

type IUserHandlers interface {
	CreateUserInDb(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	Home(w http.ResponseWriter, r *http.Request)
}
