package log

import (
	"fmt"
	"main/fe"
)

func LogInfo(args ...any) {
	fmt.Println(fe.Green, args, fe.Reset)
}
