package generator

import (
	"go/ast"
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

func processMapping(packageName string, controllerName string, mappingAnnotaion annotation.Annotaion, output *RootOutput) {
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

	route := `	app.` + method + `("/", func(c echo.Context) error {
		return c.String(200, gc000001.Index())
	})`

	output.RoutesCode = append(output.RoutesCode, route)
}
