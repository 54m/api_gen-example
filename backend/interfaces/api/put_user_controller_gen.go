// Package api ...
// generated version: 1.6.1
package api

import (
	"net/http"

	"github.com/54m/api_gen-example/backend/domain/model"
	"github.com/54m/api_gen-example/backend/domain/werror"
	"github.com/54m/api_gen-example/backend/interfaces/props"
	"github.com/54m/api_gen-example/backend/interfaces/wrapper"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
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
	var werr *werror.ErrorResponse

	ctx := c.Request().Context()

	user := &model.User{
		Age:    req.Age,
		Name:   req.Name,
		Gender: req.Gender,
	}

	if err = p.UserService.Insert(ctx, user); err != nil {
		if xerrors.As(err, &werr) {
			return nil, wrapper.NewAPIError(werr.Status, werr).SetError(err)
		}
		return nil, wrapper.NewAPIError(http.StatusInternalServerError).SetError(err)
	}

	res = &PutUserResponse{
		Status: http.StatusOK,
		User:   user,
	}

	return res, nil
}
