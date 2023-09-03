package annotation

import "strings"

type Annotaion struct {
	Name       string
	Parameters []string
}

// @Service(1, "2", 3) => []string{"1", "2", "3"}
func ParseParameters(commentText string) []string {
	parameters := []string{}

	if !strings.Contains(commentText, "(") {
		return parameters
	}

	parametersText := strings.Split(commentText, "(")[1]

	buffer := []byte{}
	for i := 0; i < len(parametersText); i++ {
		if parametersText[i] == ')' {
			if len(buffer) > 0 {
				parameters = append(parameters, string(buffer))
			}
			break
		}

		if parametersText[i] == ',' {
			if len(buffer) > 0 {
				parameters = append(parameters, string(buffer))
				buffer = []byte{}
			}

			continue
		}

		if parametersText[i] == '"' {
			continue
		}

		buffer = append(buffer, parametersText[i])
	}

	for _, parameter := range strings.Split(parametersText, ",") {
		parameters = append(parameters, strings.TrimSpace(parameter))
	}

	return parameters
}
