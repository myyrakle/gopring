package generator

import (
	"fmt"
	"go/ast"
	"path"
	"strings"

	"github.com/myyrakle/gopring/internal/annotation"
	"github.com/myyrakle/gopring/internal/ast_util"
	"github.com/myyrakle/gopring/internal/comment"
	"github.com/myyrakle/gopring/pkg/alias"
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

func processMapping(packageName string, receiverName string, mappingAnnotaion annotation.Annotaion, fn *ast.FuncDecl, originalCode *string, importsMap map[string]string, output *RootOutput) {
	functionName := fn.Name.Name

	// method 선택
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

	// Controller 어노테이션 정보 가져오기
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

	// 상세 경로 지정
	controllerPath := "/"

	if len(controllerInfo.annotation.Parameters) > 0 {
		controllerPath = controllerInfo.annotation.Parameters[0]
	}

	mappingPath := "/"

	if len(mappingAnnotaion.Parameters) > 0 {
		mappingPath = mappingAnnotaion.Parameters[0]
	}

	routePath := path.Join(controllerPath, mappingPath)

	// mapping 함수에 넘길 파라미터 목록
	parameterListToMapping := []string{}
	codeListBeforeMappingCall := []string{}

	// Parameter 목록 조회
	functionStartIndex := int(fn.Type.Params.Opening)
	startIndex := functionStartIndex
	for _, param := range fn.Type.Params.List {
		paramName := param.Names[0].Name
		selecorName := ast_util.GetSelectorNameFromType(param.Type)
		paramType := ast_util.GetTypeNameFromType(param.Type)

		paramStartIndex := int(param.Pos()) - 1
		paramEndIndex := int(param.End())

		buffer := []byte{}

		for i := startIndex; i < paramStartIndex; i++ {
			buffer = append(buffer, (*originalCode)[i])
		}
		startIndex = paramEndIndex

		// 파라미터 앞에 있으니까 아마도 주석일거임
		maybeCommentText := strings.TrimSpace(string(buffer))

		comments := comment.ParseCommentBlocks(maybeCommentText)

		if len(comments) == 0 {
			continue
		}

		for _, commentText := range comments {
			if strings.Contains(commentText, "@PathVariable") {
				annotationParameters := annotation.ParseParameters(commentText)
				pathName := ""

				if len(annotationParameters) == 0 {
					pathName = paramName
				} else {
					pathName = annotationParameters[0]
				}

				code := fmt.Sprintf(`		%s := c.Param("%s")`, pathName, pathName)

				parameterListToMapping = append(parameterListToMapping, pathName)
				codeListBeforeMappingCall = append(codeListBeforeMappingCall, code)
				continue
			}

			if strings.Contains(commentText, "@RequestParam") {
				annotationParameters := annotation.ParseParameters(commentText)
				pathName := ""

				if len(annotationParameters) == 0 {
					pathName = paramName
				} else {
					pathName = annotationParameters[0]
				}

				code := fmt.Sprintf(`		%s := c.QueryParam("%s")`, pathName, pathName)

				parameterListToMapping = append(parameterListToMapping, pathName)
				codeListBeforeMappingCall = append(codeListBeforeMappingCall, code)
				continue
			}

			if strings.Contains(commentText, "@RequestBody") {
				bodyVariableName := "body"

				typeName := ""

				if selecorName != nil {
					importPath := importsMap[*selecorName]
					packageAlias := alias.GetNextPackageAlias()

					importPackage := fmt.Sprintf("\t%s \"%s\"", packageAlias, importPath)
					output.ImportPackages[packageAlias] = importPackage
					alias.PackageAliasRefCount[packageAlias]++

					typeName = fmt.Sprintf("%s.%s", packageAlias, *paramType)
				} else {
					typeName = *paramType
				}

				code := fmt.Sprintf(`		%s := %s{}
		if err := c.Bind(body); err != nil {
			return err
		}`, bodyVariableName, typeName)

				parameterListToMapping = append(parameterListToMapping, bodyVariableName)
				codeListBeforeMappingCall = append(codeListBeforeMappingCall, code)
				continue
			}
		}
	}

	parametersToMapping := strings.Join(parameterListToMapping, ", ")
	codesBeforeMappingCall := strings.Join(codeListBeforeMappingCall, "\n")

	route := fmt.Sprintf(`	app.%s("%s", func(c echo.Context) error {
%s
		response := %s.%s(c, %s)

		return c.JSON(200, response)
	})`, method, routePath, codesBeforeMappingCall, controllerInfo.controllerAlias, functionName, parametersToMapping)

	output.RoutesCode = append(output.RoutesCode, route)
}
