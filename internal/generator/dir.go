package generator

import (
	"log"
	"os"
)

func getDirList(basePath string) []string {
	dirs, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	var dirList []string
	for _, dir := range dirs {
		if dir.IsDir() {
			dirList = append(dirList, dir.Name())
		}
	}

	return dirList
}
