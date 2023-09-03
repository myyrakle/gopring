package alias

import "fmt"

var count int = 0

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
