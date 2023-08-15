package generator

import (
	"os/exec"
	"strings"
)

func getModuleName() string {
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	moduleName := strings.TrimSpace(string(output))
	return moduleName
}

var ModuleName = getModuleName()
