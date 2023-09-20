package properties

import (
	"fmt"
	"testing"
)

func Test_ParseProperties(t *testing.T) {
	rawString := `server.port = 8888
spring.profiles = dev
spring.application.name = demo
`
	properties := ParseProperties(rawString)
	fmt.Printf("%+v\n", properties)
}
