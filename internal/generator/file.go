package generator

import (
	"go/ast"
)

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
