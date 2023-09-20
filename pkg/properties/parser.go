package properties

import "strings"

type PropertiesNode struct {
	IsLeaf bool
	Value  string
	Childs map[string]PropertiesNode
}

func ParseProperties(rawString string) PropertiesNode {
	lines := strings.Split(rawString, "\n")

	root := PropertiesNode{
		IsLeaf: false,
		Value:  "",
		Childs: make(map[string]PropertiesNode),
	}

	lastChild := root

	for _, line := range lines {
		if strings.Contains(line, "=") {
			splited := strings.Split(line, "=")
			left := splited[0]
			right := splited[1]

			leftSplited := strings.Split(left, ".")
			for _, namespace := range leftSplited {
				newChild := PropertiesNode{
					IsLeaf: false,
					Value:  "",
					Childs: make(map[string]PropertiesNode),
				}
				lastChild.Childs[namespace] = newChild
				lastChild = newChild
			}

			lastChild.IsLeaf = true
			lastChild.Value = right
		} else {
			continue
		}
	}

	return root
}
