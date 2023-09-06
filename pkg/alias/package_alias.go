package alias

import "fmt"

var count int = 0

// package alias가 참조되는 개수입니다. (0일 경우 import가 되지 않도록 할때 응용합니다.)
var PackageAliasRefCount map[string]int = make(map[string]int)

func ResetPackageCount() {
	count = 0
}

func GetNextPackageAlias() string {
	count++
	alias := fmt.Sprintf("g%06d", count)

	PackageAliasRefCount[alias] = 0
	return alias
}
