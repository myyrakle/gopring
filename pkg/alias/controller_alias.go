package alias

import "fmt"

var controllerCount int = 0

// alias가 참조되는 개수입니다. (0일 경우 import가 되지 않도록 할때 응용합니다.)
var ControllerAliasRefCount map[string]int = make(map[string]int)

var ControllerAliasMap map[string]string = make(map[string]string)

func ResetControllerCount() {
	controllerCount = 0
}

func GetNextControllerAlias() string {
	controllerCount++
	alias := fmt.Sprintf("gc%06d", controllerCount)

	ControllerAliasRefCount[alias] = 0
	return alias
}
