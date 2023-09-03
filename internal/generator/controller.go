package generator

import (
	"go/ast"
	"strings"

	"github.com/myyrakle/gopring/internal/annotation"
	"github.com/myyrakle/gopring/pkg/alias"
)

// 주석을 읽어와서 @Controller 구조체인지 검증합니다.
func getControllerAnnotation(genDecl *ast.GenDecl) *annotation.Annotaion {
	if genDecl.Doc == nil {
		return nil
	}

	if genDecl.Doc.List == nil {
		return nil
	}

	for _, comment := range genDecl.Doc.List {
		if strings.Contains(comment.Text, "@Controller") {
			parameters := annotation.ParseParameters(comment.Text)

			return &annotation.Annotaion{
				Name:       "Controller",
				Parameters: parameters,
			}
		}
	}

	return nil
}

func processContoller(packageName, structName string, structDecl *ast.StructType) string {
	var newFunctionCode string

	newFunctionCode += "func GopringNewController" + structName + "("

	for _, field := range structDecl.Fields.List {
		typeName := field.Type.(*ast.Ident).Name
		fieldName := field.Names[0].Name

		newFunctionCode += fieldName + " " + typeName + ", "
	}

	newFunctionCode += ") *" + structName + " {\n"
	newFunctionCode += "return &" + structName + "{\n"

	for _, field := range structDecl.Fields.List {
		fieldName := field.Names[0].Name

		newFunctionCode += fieldName + ": " + fieldName + ",\n"
	}

	newFunctionCode += "}\n"

	alias.PackageAliasRefCount[packageName]++

	return newFunctionCode
}
