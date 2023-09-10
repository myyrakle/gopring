package generator

import "github.com/myyrakle/gopring/internal/annotation"

type ControllerInfo struct {
	packageAlias    string
	controllerName  string
	controllerAlias string
	annotation      *annotation.Annotaion
}

var ControllerList []ControllerInfo = make([]ControllerInfo, 0)
