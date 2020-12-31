// Package user ...
// generated version: 1.6.1
package user

import (
	"net/http"

	"github.com/54m/api_gen-example/backend/domain/model"
	"github.com/54m/api_gen-example/backend/domain/werror"
	"github.com/54m/api_gen-example/backend/interfaces/props"
	"github.com/54m/api_gen-example/backend/interfaces/wrapper"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
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
	var werr *werror.ErrorResponse

	if err = c.Validate(req); err != nil {
		werr = werror.NewValidationError(err)
		return nil, wrapper.NewAPIError(werr.Status, werr).SetError(err)
	}

	ctx := c.Request().Context()

	user := &model.User{
		ID:     req.ID,
		Age:    req.Age,
		Name:   req.Name,
		Gender: req.Gender,
	}

	if err = p.UserService.Update(ctx, user); err != nil {
		if xerrors.As(err, &werr) {
			return nil, wrapper.NewAPIError(werr.Status, werr).SetError(err)
		}
		return nil, wrapper.NewAPIError(http.StatusInternalServerError).SetError(err)
	}

	res = &PatchResponse{
		Status: http.StatusOK,
		User:   user,
	}

	return res, nil
}
