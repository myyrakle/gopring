package command

import (
	"os/exec"

	"github.com/myyrakle/gopring/internal/generator"
)

func Run() {
	generator.Generate()
	exec.Command("go", "run", "dist/main.go").Run()
}
