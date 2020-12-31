// Package interfaces ...
// generated version: 1.6.1
package interfaces

import (
	"github.com/54m/api_gen-example/backend/interfaces/props"
	"github.com/labstack/echo/v4"
)

// GetHealthCheckController ...
type GetHealthCheckController struct {
	*props.ControllerProps
}

// NewGetHealthCheckController ...
func NewGetHealthCheckController(cp *props.ControllerProps) *GetHealthCheckController {
	g := &GetHealthCheckController{
		ControllerProps: cp,
	}
	return g
}

// GetHealthCheck health check api controller
// @Summary HealthCheckAPI
// @Description health check api
// @Accept json
// @Produce json
// @Success 200 {object} GetHealthCheckResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /health_check [GET]
func (g *GetHealthCheckController) GetHealthCheck(
	c echo.Context, req *GetHealthCheckRequest,
) (res *GetHealthCheckResponse, err error) {
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
