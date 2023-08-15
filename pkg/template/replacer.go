package template

import (
	"fmt"
	"strings"
)

// {{name}} 형태의 템플릿을 주어진 변수로 치환합니다.
func ReplaceTemplate(template string, templateVariables map[string]string) string {
	for key, value := range templateVariables {
		fmt.Println(key, value)
		template = strings.ReplaceAll(template, "{{"+key+"}}", value)
	}

	return template
}
