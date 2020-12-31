// Package user ...
// generated version: 1.6.1
package user

import (
	"github.com/54m/api_gen-example/backend/interfaces/props"
	"github.com/labstack/echo/v4"
)

// PatchController ...
type PatchController struct {
	*props.ControllerProps
}

// NewPatchController ...
func NewPatchController(cp *props.ControllerProps) *PatchController {
	p := &PatchController{
		ControllerProps: cp,
	}
	return p
}

// Patch user editing api controller
// @Summary EditUserAPI
// @Description user editing api
// @Accept json
// @Produce json
// @Param userID path string true "user id"
// @Param name body string false "user name"
// @Param age body integer false "user age"
// @Param gender body model.Gender false "user gender"
// @Success 200 {object} PatchResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /api/user/{userID} [PATCH]
func (p *PatchController) Patch(
	c echo.Context, req *PatchRequest,
) (res *PatchResponse, err error) {
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
