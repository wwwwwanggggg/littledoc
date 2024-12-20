package docgen

import (
	"os"
	"strings"
)

func getFileName(file *os.File) string {
	rawName := file.Name()
	var pathSlice []string
	if strings.Contains(rawName, "/") {
		pathSlice = strings.Split(rawName, "/")
	} else {
		pathSlice = strings.Split(rawName, "\\")
	}
	return strings.Split(pathSlice[len(pathSlice)-1], ".")[0]
}
