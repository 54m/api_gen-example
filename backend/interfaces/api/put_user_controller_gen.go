// Package api ...
// generated version: 1.6.1
package api

import (
	"github.com/54m/api_gen-example/backend/interfaces/props"
	"github.com/labstack/echo/v4"
)

// PutUserController ...
type PutUserController struct {
	*props.ControllerProps
}

// NewPutUserController ...
func NewPutUserController(cp *props.ControllerProps) *PutUserController {
	p := &PutUserController{
		ControllerProps: cp,
	}
	return p
}

// PutUser user creation api controller
// @Summary CreateUserAPI
// @Description user creation api
// @Accept json
// @Produce json
// @Param name body string true "user name"
// @Param age body integer true "user age"
// @Param gender body model.Gender true "user gender"
// @Success 200 {object} PutUserResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /api/user [PUT]
func (p *PutUserController) PutUser(
	c echo.Context, req *PutUserRequest,
) (res *PutUserResponse, err error) {
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
