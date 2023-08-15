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

func processFile(filename string, file *ast.File, output *RootOutput) {
	var serviceCodes []string
	var controllerCodes []string

	for _, declare := range file.Decls {
		if genDecl, ok := declare.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					structDecl, _ := typeSpec.Type.(*ast.StructType)
					structName := typeSpec.Name.Name

					if structDecl == nil {
						continue
					}

					serviceAnnotation := getServiceAnnotation(genDecl)
					if serviceAnnotation != nil {
						serviceCodes = append(serviceCodes, processService(structName, structDecl))
					}

					controllerAnnotation := getControllerAnnotation(genDecl)
					if controllerAnnotation != nil {
						controllerCodes = append(controllerCodes, processContoller(structName, structDecl))
					}
				}
			}
		}
	}
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

func processContoller(structName string, structDecl *ast.StructType) string {
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

	return newFunctionCode
}