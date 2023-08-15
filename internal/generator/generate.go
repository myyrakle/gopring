package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"

	"github.com/myyrakle/gopring/internal/templates"
	"github.com/myyrakle/gopring/pkg/template"
)

func getPackageList(basePath string) map[string]*ast.Package {
	fset := token.NewFileSet()

	packages, err := parser.ParseDir(fset, basePath, nil, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	return packages
}

func generateRecursive(basedir string, output *RootOutput) {
	packages := getPackageList(basedir)

	for packageName, asts := range packages {
		fmt.Printf(">> package [%s]...\n", packageName)

		importPackage := "\t\"" + ModuleName + "/" + strings.Replace(basedir, "src/", "dist/", 1) + "\""
		output.ImportPackages = append(output.ImportPackages, importPackage)

		for filename, file := range asts.Files {
			fmt.Printf(">> scan [%s]...\n", filename)

			text, err := os.ReadFile(filename)

			if err != nil {
				panic(err)
			}

			originalCode := string(text)
			codeToAppend := processFile(packageName, filename, file, output)

			newPath := strings.Replace(filename, "src", output.OutputBasedir, 1)

			os.WriteFile(newPath, []byte(originalCode+"\n"+codeToAppend), 0644)
		}
	}

	dirList := getDirList(basedir)

	for _, dir := range dirList {
		os.Mkdir(path.Join(output.OutputBasedir, dir), 0755)
		generateRecursive(path.Join(basedir, dir), output)
	}
}

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
