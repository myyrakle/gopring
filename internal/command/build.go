package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/myyrakle/gopring/internal/generator"
)

func Build() {
	Cleanup()

	generator.Generate()

	fmt.Println(">> go mod tidy...")
	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		panic(err)
	}
	fmt.Println(">> go mod tidy done")

	fmt.Println(">> go build dist/main.go")
	buildCommand := exec.Command("go", "build", "dist/main.go")
	buildCommand.Stdout = os.Stdout
	if err := buildCommand.Run(); err != nil {
		panic(err)
	}
	fmt.Println(">>> build success")

	Cleanup()
}
