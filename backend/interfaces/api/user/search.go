package user

import (
	"github.com/54m/api_gen-example/backend/domain/model"
	"github.com/54m/api_gen-example/backend/domain/werror"
)

// GetSearchRequest - user search request
type GetSearchRequest struct {
	Name string `json:"name" query:"name"`
	Age  int    `json:"age" query:"age"`
}

// GetSearchResponse - user search response
type GetSearchResponse struct {
	Status   int                   `json:"status"`
	User     []*model.User         `json:"payload,omitempty"`
	Messages []werror.FailedReason `json:"messages"`
}
