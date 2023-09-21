package properties

import (
	"strings"
)

type Property struct {
	Key   string
	Value string
}

type Properties = []Property

func ParseProperties(rawString string) Properties {
	lines := strings.Split(rawString, "\n")

	properties := make(Properties, 0)

	for _, line := range lines {

		if strings.Contains(line, "=") {
			splited := strings.Split(line, "=")
			left := strings.TrimSpace(splited[0])
			right := strings.TrimSpace(splited[1])

			properties = append(properties, Property{
				Key:   left,
				Value: right,
			})
		} else {
			continue
		}
	}

	return properties
}
