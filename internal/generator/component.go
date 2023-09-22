package generator

import (
	"go/ast"
	"strings"

	"github.com/myyrakle/gopring/internal/annotation"
	"github.com/myyrakle/gopring/pkg/alias"
)

// 주석을 읽어와서 @Component 계열 구조체인지 검증합니다.
func getComponentAnnotation(genDecl *ast.GenDecl) *annotation.Annotaion {
	if genDecl.Doc == nil {
		return nil
	}

	if genDecl.Doc.List == nil {
		return nil
	}

	for _, comment := range genDecl.Doc.List {
		if strings.Contains(comment.Text, "@Service") ||
			strings.Contains(comment.Text, "@Repository") ||
			strings.Contains(comment.Text, "@Component") {

			return &annotation.Annotaion{
				Name: "Component",
			}
		}
	}

	return nil
}

func processComponent(packageName string, structName string, structDecl *ast.StructType, output *RootOutput) string {
	var newFunctionCode string
	newFunctionName := "GopringNewComponent_" + structName

	output.Providers = append(output.Providers, packageName+"."+newFunctionName)

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

	return newFunctionCode
}
