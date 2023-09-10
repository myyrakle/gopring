package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/myyrakle/gopring/internal/generator"
)

func Run() {
	generator.Generate()

	fmt.Println(">> go mod tidy...")
	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		panic(err)
	}
	fmt.Println(">> go mod tidy done")

	fmt.Println(">> go run dist/main.go")
	runCommand := exec.Command("go", "run", "dist/main.go")
	runCommand.Stdout = os.Stdout
	if err := runCommand.Run(); err != nil {
		panic(err)
	}
}
