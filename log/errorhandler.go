package log

import (
	"fmt"
	"main/fe"
)

const (
	INNER_ERROR = iota
	PATH_ERROR
	READ_ERROR
	GENERATING_ERROR
)

var ErrMap = []string{
	"[INNER ERROR]",
	"[PATH ERROR]",
	"[READ ERROR]",
	"[GENERATING ERORR]",
}

func LogError(errType int, errMsg string) {
	if errType >= len(ErrMap) {
		// 不存在的错误统统改成内部错误
		errType = 0
	}
	fmt.Println(fe.Yellow, ErrMap[errType], fe.Red, errMsg, fe.Reset)
}
