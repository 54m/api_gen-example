package api

import (
	"github.com/54m/api_gen-example/backend/domain/model"
)

// PutUserRequest - user creation request
type PutUserRequest struct {
	Name   string       `json:"name" validate:"required,min=3,max=10,excludesall=!()#@{}"`
	Age    int          `json:"age" validate:"required,gt=0,lte=150"`
	Gender model.Gender `json:"gender" validate:"required,oneof=1 2 3"`
}

// PutUserResponse - user creation response
type PutUserResponse struct {
	Status int         `json:"status"`
	User   *model.User `json:"payload,omitempty"`
}
