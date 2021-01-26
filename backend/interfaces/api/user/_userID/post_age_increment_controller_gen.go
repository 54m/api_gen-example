// Package _userID ...
// generated version: 1.6.1
package _userID

import (
	"net/http"

	"github.com/54m/api_gen-example/backend/domain/werror"
	"github.com/54m/api_gen-example/backend/interfaces/props"
	"github.com/54m/api_gen-example/backend/interfaces/wrapper"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

// PostAgeIncrementController ...
type PostAgeIncrementController struct {
	*props.ControllerProps
}

// NewPostAgeIncrementController ...
func NewPostAgeIncrementController(cp *props.ControllerProps) *PostAgeIncrementController {
	p := &PostAgeIncrementController{
		ControllerProps: cp,
	}
	return p
}

// PostAgeIncrement user age increment api controller
// @Summary UserAgeIncrementAPI
// @Description user age increment api
// @Accept json
// @Produce json
// @Param userID path string true "user id"
// @Success 200 {object} PostAgeIncrementResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /api/user/{userID}/age_increment [POST]
func (p *PostAgeIncrementController) PostAgeIncrement(
	c echo.Context, req *PostAgeIncrementRequest,
) (res *PostAgeIncrementResponse, err error) {
	var werr *werror.ErrorResponse

	ctx := c.Request().Context()

	user, err := p.UserService.AgeIncrement(ctx, req.ID)
	if err != nil {
		if xerrors.As(err, &werr) {
			return nil, wrapper.NewAPIError(werr.Status, werr).SetError(err)
		}
		return nil, wrapper.NewAPIError(http.StatusInternalServerError).SetError(err)
	}

	res = &PostAgeIncrementResponse{
		Status: http.StatusOK,
		User:   user,
	}

	return res, nil
}
