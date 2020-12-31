// Package _userID ...
// generated version: 1.6.1
package _userID

import (
	"github.com/54m/api_gen-example/backend/interfaces/props"
	"github.com/labstack/echo/v4"
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
