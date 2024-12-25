package docgen

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"main/fe"
	"main/log"
	"os"
	"sync"
)

// 这个用于处理一个文件，每个文件对应于一个praser
// 此函数也会单独开协程处理
func generateDocString(file *os.File) (string, bool) {
	filename := getFileName(file)
	praser := NewParser(filename, fe.NOW+"/"+file.Name())
	wg := sync.WaitGroup{}
	stop := false
	wg.Add(2)
	reader := bufio.NewReader(file)
	go func() {
		lineCount := 0
		for {
			lineCount += 1
			line, isPerfix, err := reader.ReadLine()
			if errors.Is(err, io.EOF) {
				wg.Done()
				stop = true
				log.LogInfo(fmt.Sprintf("Finish reading %s", filename))
				// fmt.Println(fe.Green, )
				break
			} else if err != nil {
				wg.Done()
				log.LogError(log.READ_ERROR, fmt.Sprintf("Failed to read %s", filename))
				// fmt.Println(fe.Yellow, "[READ ERROR]", fe.Red)
				break
			}
			if isPerfix {
				// 表示这一行过长，应该重新读取
				log.LogError(log.READ_ERROR, fmt.Sprintf("failed to read file:%s line %d because it's too long", filename, lineCount))
				// fmt.Println(fe.Yellow, "[READ ERROR]", fe.Red, ), fe.Reset)
				continue
			}
			// fmt.Println("Queue Len", praser.QueueLen())
			praser.EnQueue(line)
		}
	}()
	for {
		if stop {
			// 如果上面已经读完了，那么在这里解析完剩下的切片
			/*
				一般prase的速度比read的速度会快一些，所以剩下的praser.QueLen()不会太大
			*/
			for i := 0; i < praser.QueueLen(); i++ {
				singleDeal(praser)
			}
			wg.Done()
			break
		}
		if praser.QueueLen() == 0 {
			// 如果队列为空，就等待
			continue
		}
		// 这里反复处理
		singleDeal(praser)
	}
	wg.Wait()
	return praser.Result(), true
}

// 单次的出队然后解析
func singleDeal(praser Parser) {
	praser.Parse()
}
