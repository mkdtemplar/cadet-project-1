package objects

import "cadet-project/models"

type GetRequest struct {
	ID uint32 `json:"id"`
}

type CreateRequest struct {
	Users *models.User `json:"users"`
}

type DeleteRequest struct {
	ID uint32 `json:"id"`
}
