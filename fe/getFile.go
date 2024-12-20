package fe

import (
	"fmt"
	"io/fs"
	"os"
)

var NOW, fullPath string

// 处理文件夹里的或者文件里的都可以
func GetFile() ([]fs.DirEntry, bool) {
	now, err := os.Getwd()
	NOW = now + "\\"
	if err != nil {
		fmt.Println(Yellow+"[PATH ERROR]"+Reset, Red, err.Error(), Reset)
		return nil, false
	}
	fullPath = now + "\\" + RePath

	res, err := os.ReadDir("./" + RePath)
	fmt.Println("repath:", "./"+RePath)
	if err != nil {
		fmt.Println(Yellow, "[PATH ERROR]", Red, err.Error(), Reset)
		return nil, false
	}
	return res, true
}
