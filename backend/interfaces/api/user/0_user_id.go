package user

import (
	"github.com/54m/api_gen-example/backend/domain/model"
	"github.com/54m/api_gen-example/backend/domain/werror"
)

// GetRequest - user acquisition request
type GetRequest struct {
	ID string `json:"userID" param:"userID" validate:"required"`
}

// GetResponse- user acquisition response
type GetResponse struct {
	Status   int                   `json:"status"`
	User     *model.User           `json:"payload,omitempty"`
	Messages []werror.FailedReason `json:"messages"`
}

// PatchRequest - user editing request
type PatchRequest struct {
	ID     string       `json:"userID" param:"userID" validate:"required"`
	Name   string       `json:"name,omitempty" validate:"min=5,max=10,excludesall=!()#@{}"`
	Age    int          `json:"age,omitempty" validate:"gt=0,lte=150"`
	Gender model.Gender `json:"gender,omitempty" validate:"oneof=0 1 2 3"`
}

// PatchResponse - user editing response
type PatchResponse struct {
	Status   int                   `json:"status"`
	User     *model.User           `json:"payload,omitempty"`
	Messages []werror.FailedReason `json:"messages"`
}

// DeleteRequest - user deletion request
type DeleteRequest struct {
	ID string `json:"userID" param:"userID" validate:"required"`
}

// DeleteResponse - user deletion response
type DeleteResponse struct {
	Status   int                   `json:"status"`
	Messages []werror.FailedReason `json:"messages"`
}
