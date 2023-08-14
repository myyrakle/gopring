package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
)

func getPackageList(basePath string) map[string]*ast.Package {
	fset := token.NewFileSet()

	packages, err := parser.ParseDir(fset, basePath, nil, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	return packages
}

func getDirList(basePath string) []string {
	dirs, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	var dirList []string
	for _, dir := range dirs {
		if dir.IsDir() {
			dirList = append(dirList, dir.Name())
		}
	}

	return dirList
}

func generateRootDefaultFile(basedir string) {
	err := os.Mkdir(basedir, 0755)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

func generateRecursive(basedir string, output *RootOutput) {
	packages := getPackageList(basedir)

	for packageName, asts := range packages {
		fmt.Printf(">> package [%s]...\n", packageName)

		for filename, _ := range asts.Files {
			fmt.Printf(">> scan [%s]...\n", filename)
		}
	}

	dirList := getDirList(basedir)

	for _, dir := range dirList {
		generateRecursive(path.Join(basedir, dir), output)
	}
}

type RootOutput struct {
	importPackages      []string
	injectedServices    []string
	injectedControllers []string
}

func generateRootFile(output *RootOutput) {
}

func Generate() {
	generateRootDefaultFile("dist")

	output := RootOutput{}
	generateRecursive("src", &output)
	generateRootFile(&output)
}
