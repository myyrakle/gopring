package generator

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

// 특정 파일에 있는 import 경로들을 /src/ => /dist/로 바꾼 전체 코드를 문자열로 반환합니다.
func replaceImportPath(originalCode string) string {
	// import 내의 경로들을 바꿔줍니다.

	os.WriteFile("_temp.go", []byte(originalCode), 0644)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "_temp.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// TODO: 제대로 문법 파싱해서 import 구문 내의 경로만 바꾸도록 수정해야 합니다.

	// import 경로에서 "src"를 "dist"로 바꿉니다.
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ImportSpec:
			if x.Path != nil {
				path := strings.Trim(x.Path.Value, `"`)
				if strings.Contains(path, "/src/") {
					newPath := strings.ReplaceAll(path, "/src/", "/dist/")
					x.Path.Value = fmt.Sprintf(`"%s"`, newPath)
				}
			}
		}
		return true
	})

	// 수정된 코드를 문자열로 변환합니다.
	var buf strings.Builder
	if err := format.Node(&buf, fset, node); err != nil {
		panic(err)
	}

	return buf.String()
}
