package docgen

import (
	"fmt"
	"main/fe"
	"main/log"
	"os"
)

func Init() {
	log.LogInfo("Begin to generate document in", fe.NOW+fe.Dest)
	// fmt.Println(fe.Green, , fe.Reset)
	for _, dir := range fe.FileDir {
		temp := "./" + fe.RePath + "/" + dir.Name()
		ef, err := os.Open(temp)
		if err != nil {
			log.LogError(log.GENERATING_ERROR, err.Error())
			// fmt.Println(fe.Yellow, "[GENERATEING ERROR]", fe.Red, err.Error(), fe.Reset)
			return

		} else {
			log.LogInfo("Generateing document for ", fe.NOW+temp)
			// fmt.Println(fe.Green,  fe.Reset)
		}
		defer ef.Close()
		res, _ := generateDocString(ef)
		fmt.Println(res)
	}
}
