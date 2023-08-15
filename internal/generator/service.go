package generator

import (
	"go/ast"
	"strings"

	"github.com/myyrakle/gopring/internal/annotation"
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

func processService(structName string, structDecl *ast.StructType) string {
	var newFunctionCode string

	newFunctionCode += "func GopringNewService" + structName + "("

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

	return newFunctionCode
}
