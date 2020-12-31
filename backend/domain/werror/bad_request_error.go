package werror

import "net/http"

// NewBadRequestError - bad request error
func NewBadRequestError(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusBadRequest,
		FailedReasons: []FailedReason{
			{
				Message: "bad request",
			},
		},
		err: err,
	}
}
