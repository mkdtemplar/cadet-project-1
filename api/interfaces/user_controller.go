package interfaces

import "net/http"

type IUserController interface {
	CreateUserInDb(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	Home(w http.ResponseWriter, r *http.Request)
}
