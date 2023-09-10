package command

import (
	"os/exec"

	"github.com/myyrakle/gopring/internal/generator"
)

func Run() {
	exec.Command("go", "mod", "tidy").Run()
	generator.Generate()
	exec.Command("go", "run", "dist/main.go").Run()
}
