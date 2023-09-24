package controller

import (
	"github.com/labstack/echo"
	"github.com/myyrakle/gopring/src/service"
)

// @Controller(/)
type HomeController struct {
	service *service.HomeService
}

// @GetMapping("/")
func (this HomeController) Index(c echo.Context) string {
	return this.service.GetHello()
}

type HealthCheckResponse struct {
	Ok bool `json:"ok"`
}

// @GetMapping("/health")
func (this *HomeController) HelathCheck(c echo.Context) HealthCheckResponse {
	return HealthCheckResponse{
		Ok: true,
	}
}

// @GetMapping("/user/:id")
func (this *HomeController) GetUserByUserId(
	c echo.Context,
	// @PathVariable("id")
	id string,
) HealthCheckResponse {
	return HealthCheckResponse{
		Ok: true,
	}
}
