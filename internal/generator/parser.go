package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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

func generateRecursive(basedir string, output *RootOutput) {
	packages := getPackageList(basedir)

	for packageName, asts := range packages {
		fmt.Printf(">> package [%s]...\n", packageName)

		for filename, file := range asts.Files {
			fmt.Printf(">> scan [%s]...\n", filename)

			processFile(filename, file, output)
		}
	}

	dirList := getDirList(basedir)

	for _, dir := range dirList {
		os.Mkdir(path.Join(output.OutputBasedir, dir), 0755)
		generateRecursive(path.Join(basedir, dir), output)
	}
}
