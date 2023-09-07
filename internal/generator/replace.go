package generator

import (
	"fmt"
	"regexp"
	"strings"
)

// 특정 파일에 있는 import 경로들을 /src/ => /dist/로 바꾼 전체 코드를 문자열로 반환합니다.
func replaceImportPath(originalCode string) string {
	// import 내의 경로들을 바꿔줍니다.

	// import가 없는 경우
	if strings.Contains(originalCode, "import") == false {
		return originalCode
	}

	// import가 소괄호 없이 있는 경우
	if strings.Contains(originalCode, "import (") == false {
		// 정규식을 사용해 import 경로를 바꿔줍니다. (import "github.com/myyrakle/gopring/src/service" => import "github.com/myyrakle/gopring/dist/service")
		regex := regexp.MustCompile(`(?m)^import\s+\"(.+)\"$`)
		newCode := regex.ReplaceAllStringFunc(string(originalCode), func(match string) string {
			path := strings.TrimPrefix(strings.TrimSuffix(match, `"`), `import "`)
			if strings.HasPrefix(path, "/src/") {
				path = strings.Replace(path, "/src/", "/dist/", 1)
			}
			return fmt.Sprintf(`import "%s"`, path)
		})

		return newCode
	}

	// import가 소괄호 로 감싸져 있는 경우
	if strings.Contains(originalCode, "import (") {
		// )가 나올때까지 반복하며 import 경로를 바꿔줍니다.

	}

	return originalCode
}
