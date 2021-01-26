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

// GetController ...
type GetController struct {
	*props.ControllerProps
}

// NewGetController ...
func NewGetController(cp *props.ControllerProps) *GetController {
	g := &GetController{
		ControllerProps: cp,
	}
	return g
}

// Get user acquisition api controller
// @Summary GetUserAPI
// @Description user acquisition api
// @Accept json
// @Produce json
// @Param userID path string true "user id"
// @Success 200 {object} GetResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /api/user/{userID} [GET]
func (g *GetController) Get(
	c echo.Context, req *GetRequest,
) (res *GetResponse, err error) {
	var werr *werror.ErrorResponse

	ctx := c.Request().Context()

	user, err := g.UserService.Get(ctx, req.ID)
	if err != nil {
		if xerrors.As(err, &werr) {
			return nil, wrapper.NewAPIError(werr.Status, werr).SetError(err)
		}
		return nil, wrapper.NewAPIError(http.StatusInternalServerError).SetError(err)
	}

	res = &GetResponse{
		Status: http.StatusOK,
		User:   user,
	}

	return res, nil
}
