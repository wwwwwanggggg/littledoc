package fe

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

const (
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Reset   = "\033[0m"
)

var RePath, RouterFileName, Dest string

var FileDir []fs.DirEntry

var RouterFile *os.File

func Init() bool {
	// 从控制台获取生成配置
	rePath := flag.String("p", "", "go file relative path")
	routerFileName := flag.String("r", "", "router file name")
	dest := flag.String("d", "doc", "doc generated's path")

	flag.Parse()

	RePath = *rePath
	RouterFileName = *routerFileName
	Dest = *dest

	file, ok := GetFile()
	// if *routerFileName != "" {
	// 	routerFile, err := os.OpenFile(*routerFileName, 1, 0666)
	// }
	// if err != nil {
	// 	fmt.Println(Yellow, "[PATH ERROR]", Reset, Red, err.Error(), NOW+RouterFileName, Reset)
	// 	return false
	// }
	// RouterFile = routerFile
	if RouterFileName != "" {
		fmt.Println(Green, "Using router file:", RouterFileName, Reset)
	}
	FileDir = file
	fmt.Println(Green, "Ready to make doc for "+fullPath, Reset)
	return ok
}
