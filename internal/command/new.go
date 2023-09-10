package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/myyrakle/gopring/internal/templates"
	"github.com/myyrakle/gopring/pkg/template"
)

func New(projectName string) {
	if err := os.Mkdir(projectName, 0755); err != nil {
		panic(err)
	}
	fmt.Println(">>> create project: " + projectName)

	modCommand := exec.Command("go", "mod", "init", projectName)
	modCommand.Dir = projectName
	modCommand.Run()
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

	os.WriteFile(projectName+"/src/service/home_service.go", []byte(templates.HOME_SERVICE), 0755)
	fmt.Println(">> create service: " + projectName + "/src/service/home_service.go")

	controllerCode := template.ReplaceTemplate(
		templates.HOME_CONTROLLER,
		map[string]string{
			"projectName": projectName,
		},
	)
	os.WriteFile(projectName+"/src/controller/home_controller.go", []byte(controllerCode), 0755)
	fmt.Println(">> create controller: " + projectName + "/src/controller/home_controller.go")

	fmt.Println(">>> finished")
}
