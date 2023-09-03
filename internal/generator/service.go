package generator

import (
	"go/ast"
	"strings"

	"github.com/myyrakle/gopring/internal/annotation"
	"github.com/myyrakle/gopring/pkg/alias"
)

// 주석을 읽어와서 @Service 구조체인지 검증합니다.
func getServiceAnnotation(genDecl *ast.GenDecl) *annotation.Annotaion {
	if genDecl.Doc == nil {
		return nil
	}

	if genDecl.Doc.List == nil {
		return nil
	}

	for _, comment := range genDecl.Doc.List {
		if strings.Contains(comment.Text, "@Service") {

			return &annotation.Annotaion{
				Name: "Service",
			}
		}
	}

	return nil
}

func processService(packageName string, structName string, structDecl *ast.StructType, output *RootOutput) string {
	var newFunctionCode string
	newFunctionName := "GopringNewService" + structName

	output.Providers = append(output.Providers, packageName+"."+newFunctionName)

	newFunctionCode += "func " + newFunctionName + "("

	for _, field := range structDecl.Fields.List {
		typeName := field.Type.(*ast.Ident).Name
		fieldName := field.Names[0].Name

		newFunctionCode += fieldName + " " + typeName + ", "
	}

	newFunctionCode += ") *" + structName + " {\n"
	newFunctionCode += "\treturn &" + structName + "{\n"

	for _, field := range structDecl.Fields.List {
		fieldName := field.Names[0].Name

		newFunctionCode += fieldName + ": " + fieldName + ",\n"
	}

	newFunctionCode += "\t}\n"
	newFunctionCode += "}\n"

	alias.PackageAliasRefCount[packageName]++

	return newFunctionCode
}
