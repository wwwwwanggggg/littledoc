package docgen

import (
	"container/list"
	"fmt"
	"main/log"
)

// 文件解释器,现在是单行单行的解释了
type Parser struct {
	isMutilLine bool       // 是否在处理多行注释
	queue       *list.List // 正在等待处理的行队列
	res         string     // 产生的结果字符串
	dest        string     // 结果文件的名称
	count       int        // 当前行的行数
	nowPath     string     // 当前文件的绝对路径
}

func NewParser(filename string, nowPath string) Parser {
	return Parser{
		isMutilLine: false,
		queue:       list.New(),
		res:         "",
		dest:        filename,
		nowPath:     nowPath,
	}
}

// 解释单行
func (p *Parser) Parse() {
	if !p.isMutilLine {
		line := p.DeQueue()
		res, ok := matchAndDeal(line)
		if !ok {
			log.LogError(log.CONTENT_ERROR, fmt.Sprintf("Unknown tag in %s line %d", p.nowPath, p.count))
			return
		}
		p.res += res + "\n"
	}
	p.count += 1
}

// 得到结果
func (p *Parser) Result() string {
	return p.res
}

// 入队
func (p *Parser) EnQueue(arg any) {
	p.queue.PushBack(arg)
}

// 出队
func (p *Parser) DeQueue() string {
	elem := p.queue.Front()
	p.queue.Remove(elem)
	value, _ := elem.Value.(string)
	return value
}

// 获取队列长度
func (p *Parser) QueueLen() int {
	return p.queue.Len()
}
