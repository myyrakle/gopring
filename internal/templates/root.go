package templates

const ROOT_CODE_TEMPLATE = `package main

import (
	"os"

	echo "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	fx "go.uber.org/fx" 
	"github.com/myyrakle/gopring/pkg/properties"
{{importPackages}}
)

func RunGopring(appProperties *properties.Properties, {{params}}) {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

{{routes}}

	port := ":8080"

	if appProperties != nil {
		property := properties.FindByKey(*appProperties, "server.port")
		if property != nil {
			port = ":" + property.Value
		}
	}

	app.Logger.Fatal(app.Start(port))

	app.Logger.Fatal(app.Start(":8080"))
}

func LoadApplicationProperties() *properties.Properties {
	if _, err := os.Stat("application.properties"); os.IsNotExist(err) {
		return &properties.Properties{}
	} else {
		if fileData, err := os.ReadFile("application.properties"); err != nil {
			panic(err)
		} else {
			parsedProperties := properties.ParseProperties(string(fileData))
			return &parsedProperties
		}
	}
}

func main() {
	providers := fx.Provide(
		LoadApplicationProperties,
		{{providers}}
	)

	fx.New(providers, fx.Invoke(RunGopring)).Run()
}
`

const APPLICATION_PROPERTIES_TEMPLATE = `server.port = 7777
spring.profiles = dev
spring.application.name = demo`
