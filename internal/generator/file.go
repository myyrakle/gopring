package generator

import (
	"go/ast"
	"strings"
)

// 각 파일 하나하나에 대한 처리를 수행합니다.
func processFile(packageName string, filename string, file *ast.File, originalCode *string, originalBytes []byte, output *RootOutput) string {
	var componentCodes []string
	var controllerCodes []string

	importsMap := make(map[string]string)
	for _, importSpec := range file.Imports {
		importPath := strings.Replace(strings.ReplaceAll(importSpec.Path.Value, "\"", ""), "/src/", "/dist/", 1)
		pkgName := ""

		if importSpec.Name != nil {
			pkgName = importSpec.Name.Name
		} else {
			splitedPath := strings.Split(importPath, "/")
			pkgName = splitedPath[len(splitedPath)-1]
		}
		importsMap[pkgName] = importPath
	}

	// Controller & Service
	for _, declare := range file.Decls {
		if genDecl, ok := declare.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					structDecl, _ := typeSpec.Type.(*ast.StructType)
					structName := typeSpec.Name.Name

					if structDecl == nil {
						continue
					}

					componentAnnotation := getComponentAnnotation(genDecl)
					if componentAnnotation != nil {
						componentCodes = append(componentCodes, processComponent(packageName, structName, structDecl, output))
					}

					controllerAnnotation := getControllerAnnotation(genDecl)
					if controllerAnnotation != nil {
						controllerCodes = append(controllerCodes, processContoller(packageName, *controllerAnnotation, structName, structDecl, output))
					}
				}
			}
		}
	}

	// Mappings
	for _, declare := range file.Decls {
		if fn, ok := declare.(*ast.FuncDecl); ok {
			mappingAnnotaion := getMappingAnnotaion(fn)
			receiverName := ""

			if mappingAnnotaion != nil {
				if fn.Recv != nil {
					if len(fn.Recv.List) > 0 {

						// 포인터 리시버
						if starExpr, ok := fn.Recv.List[0].Type.(*ast.StarExpr); ok {
							receiverName = starExpr.X.(*ast.Ident).Name
						}

						// 값 리시버
						if ident, ok := fn.Recv.List[0].Type.(*ast.Ident); ok {
							receiverName = ident.Name
						}
					}
				}

			}

			if receiverName != "" {
				processMapping(packageName, receiverName, *mappingAnnotaion, fn, originalBytes, importsMap, output)
			}
		}
	}

	codeToAppend := ""

	// 서비스 코드들을 apppend할 코드에 추가합니다.
	for _, componentCode := range componentCodes {
		codeToAppend += componentCode + "\n"
	}

	codeToAppend += "\n"

	for _, controllerCode := range controllerCodes {
		codeToAppend += controllerCode + "\n"
	}

	return codeToAppend
}
