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

// 处理单行注释
func matchAndDeal(content string) (string, bool) {
	res := ""
	if !strings.HasPrefix(content, "// @") {
		return res, true
	}
	tag := strings.Split(content[4:], " ")[0]
	switch tag {
	case "method":
		{
			res += "#### 请求方法" + content[10:] + "\n\n"
		}
	case "router":
		{
			res += "#### 请求URL" + content[10:] + "\n\n"
		}
	case "type":
		{
			res += "#### 传入格式" + content[8:] + "\n\n"
		}
	default:
		{
			return "", false
		}
	}
	return res, true
}
