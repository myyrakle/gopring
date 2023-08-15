package generator

import (
	"fmt"
	"os"

	"github.com/myyrakle/gopring/internal/templates"
	"github.com/myyrakle/gopring/pkg/template"
)

func generateRootDefaultFile(basedir string) {
	err := os.Mkdir(basedir, 0755)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

type RootOutput struct {
	OutputBasedir       string
	ImportPackages      []string
	Providers           []string
	InjectedServices    []string
	InjectedControllers []string
}

func generateRootFile(output *RootOutput) {
	fmt.Printf(">> generate root file...\n")

	importPackages := ""
	for _, importPackage := range output.ImportPackages {
		importPackages += importPackage + "\n"
	}

	providers := ""
	for _, provider := range output.Providers {
		providers += provider + ",\n"
	}

	templateMap := map[string]string{"importPackages": importPackages, "providers": providers}

	code := template.ReplaceTemplate(templates.ROOT_CODE_TEMPLATE, templateMap)

	os.WriteFile("dist/main.go", []byte(code), 0644)
}

func Generate() {
	generateRootDefaultFile("dist")

	output := RootOutput{
		OutputBasedir: "dist",
	}
	generateRecursive("src", &output)
	generateRootFile(&output)
}