// Package _userID ...
// generated version: 1.6.1
package _userID

import (
	"github.com/54m/api_gen-example/backend/interfaces/props"
	"github.com/labstack/echo/v4"
)

// PatchAgeIncrementController ...
type PatchAgeIncrementController struct {
	*props.ControllerProps
}

// NewPatchAgeIncrementController ...
func NewPatchAgeIncrementController(cp *props.ControllerProps) *PatchAgeIncrementController {
	p := &PatchAgeIncrementController{
		ControllerProps: cp,
	}
	return p
}

// PatchAgeIncrement ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string true "user id"
// @Success 200 {object} PatchAgeIncrementResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /api/user/{userID}/age_increment [PATCH]
func (p *PatchAgeIncrementController) PatchAgeIncrement(
	c echo.Context, req *PatchAgeIncrementRequest,
) (res *PatchAgeIncrementResponse, err error) {
	// API Error Usage: github.com/54m/api_gen-example/backend/interfaces/wrapper
	//
	// return nil, wrapper.NewAPIError(http.StatusBadRequest)
	//
	// return nil, wrapper.NewAPIError(http.StatusBadRequest).SetError(err)
	//
	// body := map[string]interface{}{
	// 	"code": http.StatusBadRequest,
	// 	"message": "invalid request parameter.",
	// }
	// return nil, wrapper.NewAPIError(http.StatusBadRequest, body).SetError(err)
	panic("require implements.") // FIXME require implements.
}
