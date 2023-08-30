package generate_alias

import "fmt"

var count int = 0

func ResetImportCount() {
	count = 0
}

func GetNextImportAlias() string {
	count++
	return fmt.Sprintf("%06d", count)
}
