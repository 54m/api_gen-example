package api

import (
	"github.com/54m/api_gen-example/backend/domain/model"
)

// PutUserRequest - user creation request
type PutUserRequest struct {
	Name   string       `json:"name"`
	Age    int          `json:"age"`
	Gender model.Gender `json:"gender"`
}

// PutUserResponse - user creation response
type PutUserResponse struct {
	Status int         `json:"status"`
	User   *model.User `json:"payload,omitempty"`
}
