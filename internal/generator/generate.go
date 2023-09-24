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
	"github.com/myyrakle/gopring/pkg/alias"
	"github.com/myyrakle/gopring/pkg/template"
)

// 해당 경로의 패키지 목록을 가져옵니다.
func getPackageList(basePath string) map[string]*ast.Package {
	fset := token.NewFileSet()

	packages, err := parser.ParseDir(fset, basePath, nil, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	return packages
}

// 재귀적으로 해당 경로와 하위 경로의 디렉토리 목록을 조회합니다.
func generateRecursive(basedir string, output *RootOutput) {
	packages := getPackageList(basedir)
	packageAlias := alias.GetNextPackageAlias()

	for packageName, asts := range packages {
		fmt.Printf(">> package [%s]...\n", packageName)

		importPackage := "\t" + packageAlias + " \"" + ModuleName + "/" + strings.Replace(basedir, "src/", "dist/", 1) + "\""
		output.ImportPackages[packageAlias] = importPackage

		for filename, _ := range asts.Files {
			fset := token.NewFileSet()

			fmt.Printf(">> scan [%s]...\n", filename)
			file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)

			if err != nil {
				panic(err)
			}

			originalBytes, err := os.ReadFile(filename)

			if err != nil {
				panic(err)
			}

			originalCode := string(originalBytes)
			replacedCode := replaceImportPath(originalCode)

			codeToAppend := processFile(packageAlias, filename, file, &originalCode, originalBytes, output)

			newPath := strings.Replace(filename, "src", output.OutputBasedir, 1)

			os.WriteFile(newPath, []byte(replacedCode+"\n"+codeToAppend), 0644)
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

func generateRootFile(output *RootOutput) {
	fmt.Printf(">> generate root file...\n")

	importPackages := ""
	for packageAlias, importPackage := range output.ImportPackages {
		if alias.PackageAliasRefCount[packageAlias] > 0 {
			importPackages += importPackage + "\n"
		}
	}

	providers := ""
	for _, provider := range output.Providers {
		providers += provider + ",\n\t\t"
	}

	routes := ""
	for _, route := range output.RoutesCode {
		routes += route + "\n"
	}

	params := ""
	for _, param := range output.RunGopringParameters {
		params += param + ", "
	}

	templateMap := map[string]string{
		"importPackages": importPackages,
		"providers":      providers,
		"routes":         routes,
		"params":         params,
	}

	code := template.ReplaceTemplate(templates.ROOT_CODE_TEMPLATE, templateMap)

	os.WriteFile("dist/main.go", []byte(code), 0644)
}

func Generate() {
	generateRootDefaultFile("dist")

	output := RootOutput{
		OutputBasedir:  "dist",
		ImportPackages: map[string]string{},
	}
	generateRecursive("src", &output)
	generateRootFile(&output)

	cleanup()
}

func cleanup() {
	if os.IsExist(os.Remove("_temp.go")) {
		os.Remove("_temp.go")
	}
}
