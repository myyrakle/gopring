package alias

import "fmt"

var count int = 0

func ResetPackageCount() {
	count = 0
}

func GetNextPackageAlias() string {
	count++
	return fmt.Sprintf("g%06d", count)
}
