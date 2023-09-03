package generator

import (
	"go/ast"
)

// 각 파일 하나하나에 대한 처리를 수행합니다.
func processFile(packageName string, filename string, file *ast.File, output *RootOutput) string {
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
						serviceCodes = append(serviceCodes, processService(packageName, structName, structDecl, output))
					}

					controllerAnnotation := getControllerAnnotation(genDecl)
					if controllerAnnotation != nil {
						controllerCodes = append(controllerCodes, processContoller(packageName, *controllerAnnotation, structName, structDecl))
					}
				}
			}
		}
	}

	codeToAppend := ""

	// 서비스 코드들을 apppend할 코드에 추가합니다.
	for _, serviceCode := range serviceCodes {
		codeToAppend += serviceCode + "\n"
	}

	codeToAppend += "\n"

	for _, controllerCode := range controllerCodes {
		codeToAppend += controllerCode + "\n"
	}

	//fmt.Println(codeToAppend)

	return codeToAppend
}
