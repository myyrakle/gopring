package command

import (
	"os"
)

func Cleanup() {
	if os.IsExist(os.Remove("_temp.go")) {
		os.Remove("_temp.go")
	}

	if os.IsExist(os.RemoveAll("dist")) {
		os.RemoveAll("dist")
	}
}
