package _userID

import (
	"github.com/54m/api_gen-example/backend/domain/model"
	"github.com/54m/api_gen-example/backend/domain/werror"
)

type PostAgeIncrementRequest struct {
	ID string `json:"userID" param:"userID" validate:"required"`
}

type PostAgeIncrementResponse struct {
	Status   int                   `json:"status"`
	User     *model.User           `json:"payload,omitempty"`
	Messages []werror.FailedReason `json:"messages"`
}
