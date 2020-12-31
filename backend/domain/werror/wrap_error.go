package werror

import (
	"fmt"

	"golang.org/x/xerrors"
)

// ErrorResponse - error response
type ErrorResponse struct {
	Status        int            `json:"status"`
	FailedReasons []FailedReason `json:"messages"`
	err           error
}

// interface check
var (
	_ error           = new(ErrorResponse)
	_ xerrors.Wrapper = new(ErrorResponse)
)

// Error - to string
func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("error response: status=%d, internal=%+v", err.Status, err.err)
}

// Unwrap - unwrap error
func (err *ErrorResponse) Unwrap() error {
	return err.err
}

// FailedReason - failed reason
type FailedReason struct {
	Message string `json:"message"`
}
