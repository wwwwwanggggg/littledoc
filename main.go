package main

import (
	"fmt"
	"main/docgen"
	"main/fe"
)

func main() {
	fmt.Println(fe.Yellow+"[BEGIN]", fe.Reset)
	defer func() {
		fmt.Println(fe.Yellow+"[END]", fe.Reset)
	}()
	if ok := fe.Init(); !ok {
		return
	}
	docgen.Init()
}
