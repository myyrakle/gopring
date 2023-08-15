package annotation

import "strings"

type Annotaion struct {
	Name       string
	Parameters []string
}

// "@Service(1, 2, 3)" => []string{"1", "2", "3"}
func ParseParameters(commentText string) []string {
	parameters := []string{}

	if !strings.Contains(commentText, "(") {
		return parameters
	}

	parametersText := strings.TrimSuffix(strings.Split(commentText, "(")[1], ")")

	for _, parameter := range strings.Split(parametersText, ",") {
		parameters = append(parameters, strings.TrimSpace(parameter))
	}

	return parameters
}
