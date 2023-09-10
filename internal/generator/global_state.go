package generator

import "github.com/myyrakle/gopring/internal/annotation"

type ControllerInfo struct {
	packageAlias    string
	controllerName  string
	controllerAlias string
	annotation      *annotation.Annotaion
}

var ControllerList []ControllerInfo = make([]ControllerInfo, 0)

func FindByPackageAliasAndControllerName(packageAlias string, controllerName string) *ControllerInfo {
	for _, controller := range ControllerList {
		if controller.packageAlias == packageAlias && controller.controllerName == controllerName {
			return &controller
		}
	}

	return nil
}
