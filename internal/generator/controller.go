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

func processContoller(packageName string, annotaion annotation.Annotaion, structName string, structDecl *ast.StructType, output *RootOutput) string {
	var newFunctionCode string
	newFunctionName := "GopringNewController" + structName

	controllerAlias := alias.GetNextControllerAlias()

	output.Providers = append(output.Providers, packageName+"."+newFunctionName)
	output.RunGopringParameters = append(output.RunGopringParameters, controllerAlias+" "+packageName+"."+structName)

	newFunctionCode += "func " + newFunctionName + "("

	for _, field := range structDecl.Fields.List {
		fieldName := field.Names[0].Name
		typeExpr := field.Type
		typeName := ""

		if expr, ok := typeExpr.(*ast.StarExpr); ok {
			typeName += "*"
			typeExpr = expr.X
		}

		if expr, ok := typeExpr.(*ast.SelectorExpr); ok {
			if ident, ok := expr.X.(*ast.Ident); ok {
				packageName := ident.Name
				typeName += packageName + "." + expr.Sel.Name

				field := fieldName + " " + typeName

				newFunctionCode += field + ", "
				continue
			}
		}

		if ident, ok := typeExpr.(*ast.Ident); ok {
			typeName := ident.Name

			field := fieldName + " " + typeName

			newFunctionCode += field + ", "
			continue
		}
	}

	newFunctionCode += ") *" + structName + " {\n"
	newFunctionCode += "\treturn &" + structName + "{\n"

	for _, field := range structDecl.Fields.List {
		fieldName := field.Names[0].Name

		newFunctionCode += "\t\t" + fieldName + ": " + fieldName + ",\n"
	}

	newFunctionCode += "\t}\n"
	newFunctionCode += "}\n"

	alias.PackageAliasRefCount[packageName]++

	controllerInfo := ControllerInfo{
		controllerName:  structName,
		controllerAlias: controllerAlias,
		packageAlias:    packageName,
		annotation:      &annotaion,
	}
	ControllerList = append(ControllerList, controllerInfo)

	return newFunctionCode
}
