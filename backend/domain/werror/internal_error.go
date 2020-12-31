package werror

import "net/http"

// NewInternalError - internal server error
func NewInternalError(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusInternalServerError,
		FailedReasons: []FailedReason{
			{
				Message: "internal server error",
			},
		},
		err: err,
	}
}
