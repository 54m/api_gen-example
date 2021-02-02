package werror

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// NewValidationError - validation error
func NewValidationError(er error) *ErrorResponse {
	errs := er.(validator.ValidationErrors)
	reasons := make([]FailedReason, len(errs))
	for i, err := range errs {
		reasons[i].Message = fmt.Sprintf("%s: %s", err.Field(), err.ActualTag())
	}
	return &ErrorResponse{
		Status:        http.StatusUnprocessableEntity,
		FailedReasons: reasons,
		err:           er,
	}
}
