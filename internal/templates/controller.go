package templates

const HOME_CONTROLLER = `package controller

import (
	"github.com/labstack/echo"
	"{{projectName}}/src/service"
)

// @Controller(/)
type HomeController struct {
	service *service.HomeService
}

// @GetMapping("/")
func (this HomeController) Index(c echo.Context) string {
	return this.service.GetHello()
}

// @GetMapping("/health")
func (this *HomeController) HelathCheck(c echo.Context) string {
	return this.service.GetHello()
}`
