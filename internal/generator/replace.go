package generator

import (
	"strings"
)

// 특정 파일에 있는 import 경로들을 /src/ => /dist/로 바꾼 전체 코드를 문자열로 반환합니다.
func replaceImportPath(originalCode string) string {
	// import 내의 경로들을 바꿔줍니다.

	// TODO: 제대로 문법 파싱해서 import 구문 내의 경로만 바꾸도록 수정해야 합니다.

	return strings.ReplaceAll(originalCode, "/src/", "/dist/")
}
