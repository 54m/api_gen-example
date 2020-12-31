// Package user ...
// generated version: 1.6.1
package user

import (
	"github.com/54m/api_gen-example/backend/interfaces/props"
	"github.com/labstack/echo/v4"
)

// GetSearchController ...
type GetSearchController struct {
	*props.ControllerProps
}

// NewGetSearchController ...
func NewGetSearchController(cp *props.ControllerProps) *GetSearchController {
	g := &GetSearchController{
		ControllerProps: cp,
	}
	return g
}

// GetSearch user search api controller
// @Summary SearchUserAPI
// @Description user search api
// @Accept json
// @Produce json
// @Param name query string false "user name"
// @Param age query integer false "user age"
// @Param gender query model.Gender false "user gender"
// @Success 200 {object} GetSearchResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /api/user/search [GET]
func (g *GetSearchController) GetSearch(
	c echo.Context, req *GetSearchRequest,
) (res *GetSearchResponse, err error) {
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
