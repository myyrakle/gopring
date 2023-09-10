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

	if len(genDecl.Doc.List) == 0 {
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

func processContoller(packageName string, annotaion annotation.Annotaion, structName string, structDecl *ast.StructType) string {
	var newFunctionCode string

	newFunctionCode += "func GopringNewController" + structName + "("

	for _, field := range structDecl.Fields.List {
		if ident, ok := field.Type.(*ast.Ident); ok {
			typeName := ident.Name
			fieldName := field.Names[0].Name

			newFunctionCode += fieldName + " " + typeName + ", "
		}
	}

	newFunctionCode += ") *" + structName + " {\n"
	newFunctionCode += "return &" + structName + "{\n"

	for _, field := range structDecl.Fields.List {
		fieldName := field.Names[0].Name

		newFunctionCode += fieldName + ": " + fieldName + ",\n"
	}

	newFunctionCode += "\t}\n"
	newFunctionCode += "}\n"

	alias.PackageAliasRefCount[packageName]++

	return newFunctionCode
}

func getMappingAnnotaion(ast *ast.FuncDecl) *annotation.Annotaion {
	if ast == nil {
		return nil
	}

	if ast.Doc == nil {
		return nil
	}

	if ast.Doc.List == nil {
		return nil
	}

	if len(ast.Doc.List) == 0 {
		return nil
	}

	for _, comment := range ast.Doc.List {
		if strings.Contains(comment.Text, "@GetMapping") {

			parameters := annotation.ParseParameters(comment.Text)

			return &annotation.Annotaion{
				Name:       "GetMapping",
				Parameters: parameters,
			}
		}

		if strings.Contains(comment.Text, "@PostMapping") {
			parameters := annotation.ParseParameters(comment.Text)

			return &annotation.Annotaion{
				Name:       "PostMapping",
				Parameters: parameters,
			}
		}

		if strings.Contains(comment.Text, "@PutMapping") {
			parameters := annotation.ParseParameters(comment.Text)

			return &annotation.Annotaion{
				Name:       "PutMapping",
				Parameters: parameters,
			}
		}

		if strings.Contains(comment.Text, "@DeleteMapping") {
			parameters := annotation.ParseParameters(comment.Text)

			return &annotation.Annotaion{
				Name:       "DeleteMapping",
				Parameters: parameters,
			}
		}

		if strings.Contains(comment.Text, "@PatchMapping") {
			parameters := annotation.ParseParameters(comment.Text)

			return &annotation.Annotaion{
				Name:       "PatchMapping",
				Parameters: parameters,
			}
		}

		if strings.Contains(comment.Text, "@HeadMapping") {
			parameters := annotation.ParseParameters(comment.Text)

			return &annotation.Annotaion{
				Name:       "HeadMapping",
				Parameters: parameters,
			}
		}

		if strings.Contains(comment.Text, "@OptionsMapping") {
			parameters := annotation.ParseParameters(comment.Text)

			return &annotation.Annotaion{
				Name:       "OptionsMapping",
				Parameters: parameters,
			}
		}
	}

	return nil
}
