package user

import (
	"github.com/54m/api_gen-example/backend/domain/model"
	"github.com/54m/api_gen-example/backend/domain/werror"
)

// GetRequest - user acquisition request
type GetRequest struct {
	ID string `json:"userID" param:"userID"`
}

// GetResponse- user acquisition response
type GetResponse struct {
	Status   int                   `json:"status"`
	User     *model.User           `json:"payload,omitempty"`
	Messages []werror.FailedReason `json:"messages"`
}

// PatchRequest - user editing request
type PatchRequest struct {
	ID     string       `json:"userID" param:"userID"`
	Name   string       `json:"name,omitempty"`
	Age    int          `json:"age,omitempty"`
	Gender model.Gender `json:"gender,omitempty"`
}

// PatchResponse - user editing response
type PatchResponse struct {
	Status   int                   `json:"status"`
	User     *model.User           `json:"payload,omitempty"`
	Messages []werror.FailedReason `json:"messages"`
}

// DeleteRequest - user deletion request
type DeleteRequest struct {
	ID string `json:"userID" param:"userID"`
}

// DeleteResponse - user deletion response
type DeleteResponse struct {
	Status   int                   `json:"status"`
	Messages []werror.FailedReason `json:"messages"`
}
