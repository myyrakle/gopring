package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/myyrakle/gopring/internal/templates"
)

func New(projectName string) {
	if err := os.Mkdir(projectName, 0755); err != nil {
		panic(err)
	}
	fmt.Println(">>> create project: " + projectName)

	exec.Command("go", "mod", "init", projectName).Run()
	fmt.Println(">> go mod init " + projectName)

	if err := os.Mkdir(projectName+"/src", 0755); err != nil {
		panic(err)
	}
	fmt.Println(">> create " + projectName + "/src")

	if err := os.Mkdir(projectName+"/src/controller", 0755); err != nil {
		panic(err)
	}
	fmt.Println(">> create " + projectName + "/src/controller")

	if err := os.Mkdir(projectName+"/src/service", 0755); err != nil {
		panic(err)
	}
	fmt.Println(">> create " + projectName + "/src/service")

	os.WriteFile(projectName+"/src/service/home_service", []byte(templates.HOME_SERVICE), 0755)
	fmt.Println(">> create service: " + projectName + "/src/service/home_service")

	os.WriteFile(projectName+"/src/controller/home_controller", []byte(templates.HOME_CONTROLLER), 0755)
	fmt.Println(">> create controller: " + projectName + "/src/controller/home_controller")

	fmt.Println(">>> finished")
}
