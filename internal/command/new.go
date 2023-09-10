package command

import (
	"os"
	"os/exec"
)

func New(projectName string) {
	if err := os.Mkdir(projectName, 0755); err != nil {
		panic(err)
	}

	exec.Command("go", "mod", "init", projectName).Run()

	if err := os.Mkdir(projectName+"/src", 0755); err != nil {
		panic(err)
	}

	
}
