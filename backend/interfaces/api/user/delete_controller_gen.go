// Package user ...
// generated version: 1.6.1
package user

import (
	"net/http"

	"github.com/54m/api_gen-example/backend/domain/werror"
	"github.com/54m/api_gen-example/backend/interfaces/props"
	"github.com/54m/api_gen-example/backend/interfaces/wrapper"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
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
	var werr *werror.ErrorResponse

	if err = c.Validate(req); err != nil {
		werr = werror.NewValidationError(err)
		return nil, wrapper.NewAPIError(werr.Status, werr).SetError(err)
	}

	ctx := c.Request().Context()

	if err = d.UserService.DeleteByID(ctx, req.ID); err != nil {
		if xerrors.As(err, &werr) {
			return nil, wrapper.NewAPIError(werr.Status, werr).SetError(err)
		}
		return nil, wrapper.NewAPIError(http.StatusInternalServerError).SetError(err)
	}

	res = &DeleteResponse{
		Status: http.StatusOK,
	}

	return res, nil
}
