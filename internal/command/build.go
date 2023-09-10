package command

import (
	"os/exec"

	"github.com/myyrakle/gopring/internal/generator"
)

func Build() {
	exec.Command("go", "mod", "tidy").Run()
	generator.Generate()
}
