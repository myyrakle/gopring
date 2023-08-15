package templates

const ROOT_CODE_TEMPLATE = `package main

package main

import (
	echo "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	fx "go.uber.org/fx"
)

func RunGopring() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Logger.Fatal(app.Start(":8080"))
}

func main() {
	providers := fx.Provide({{providers}})

	fx.New(providers, fx.Invoke(RunGopring)).Run()
}
`
