package interfaces

import "net/http"

type ILoginController interface {
	Login(w http.ResponseWriter, r *http.Request)
}
