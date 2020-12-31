// Package user ...
// generated version: 1.6.1
package user

import (
	"github.com/54m/api_gen-example/backend/interfaces/props"
	"github.com/labstack/echo/v4"
)

// DeleteController ...
type DeleteController struct {
	*props.ControllerProps
}

// NewDeleteController ...
func NewDeleteController(cp *props.ControllerProps) *DeleteController {
	d := &DeleteController{
		ControllerProps: cp,
	}
	return d
}

// Delete user deletion controller
// @Summary DeleteUserAPI
// @Description user deletion api
// @Accept json
// @Produce json
// @Param userID path string true "user id"
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /api/user/{userID} [DELETE]
func (d *DeleteController) Delete(
	c echo.Context, req *DeleteRequest,
) (res *DeleteResponse, err error) {
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
