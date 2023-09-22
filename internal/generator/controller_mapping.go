package generator

import (
	"fmt"
	"go/ast"
	"path"
	"strings"

	"github.com/myyrakle/gopring/internal/annotation"
)

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

func processMapping(packageName string, receiverName string, mappingAnnotaion annotation.Annotaion, fn *ast.FuncDecl, output *RootOutput) {
	functionName := fn.Name.Name

	method := ""

	switch mappingAnnotaion.Name {
	case "GetMapping":
		method = "GET"
	case "PostMapping":
		method = "POST"
	case "PutMapping":
		method = "PUT"
	case "DeleteMapping":
		method = "DELETE"
	case "PatchMapping":
		method = "PATCH"
	case "HeadMapping":
		method = "HEAD"
	case "OptionsMapping":
		method = "OPTIONS"
	default:
		return
	}

	controllerInfo := FindByPackageAliasAndControllerName(packageName, receiverName)

	if controllerInfo == nil {
		return
	}

	if controllerInfo.annotation == nil {
		return
	}

	if controllerInfo.annotation.Parameters == nil {
		return
	}

	controllerPath := "/"

	if len(controllerInfo.annotation.Parameters) > 0 {
		controllerPath = controllerInfo.annotation.Parameters[0]
	}

	mappingPath := "/"

	if len(mappingAnnotaion.Parameters) > 0 {
		mappingPath = mappingAnnotaion.Parameters[0]
	}

	routePath := path.Join(controllerPath, mappingPath)

	route := fmt.Sprintf(`	app.%s("%s", func(c echo.Context) error {
		response := %s.%s(c)

		return c.JSON(200, response)
	})`, method, routePath, controllerInfo.controllerAlias, functionName)

	output.RoutesCode = append(output.RoutesCode, route)
}
