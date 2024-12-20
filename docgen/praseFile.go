package docgen

import (
	"bufio"
	"container/list"
	"fmt"
	"main/fe"
	"main/log"
	"os"
	"strings"
)

type Parser struct {
	multiLine bool       // 是否在处理多行注释
	res       string     // 某个文件的解析结果
	lines     *[]string  // 正在处理的多行
	queue     *list.List // 正在处理的文件切片队列
	length    int        // 正在处理的切片总行数
	now       int        // 正在处理的行数
	dest      string     // 写入的文件的文件名 不包含md后缀
}

// 返回未经设置的parser
func NewParser(dest string) Parser {
	parser := Parser{
		lines:     nil,
		length:    0,
		now:       0,
		multiLine: false,
		res:       "",
		queue:     list.New(),
		dest:      dest,
	}
	return parser
}

// 入队
func (p *Parser) EnQueue(item any) *list.Element {
	return p.queue.PushBack(item)
}

// 出队
func (p *Parser) DeQueue() *list.Element {
	item := p.queue.Front()
	p.queue.Remove(item)
	return item
}

// 返回队列长
func (p *Parser) QueueLen() int {
	return p.queue.Len()
}

// 返回处理结果
func (p Parser) Result() string {
	return p.res
}

// 更新正在处理的切片
func (p *Parser) Update() bool {
	item := p.DeQueue()
	content, ok := item.Value.(string)
	if !ok {
		return false
	}
	temp := strings.Split(content, "\n")
	p.lines = &temp
	p.length = len(*p.lines)
	p.now = 0
	return true
}

// 解析文件
func (p *Parser) Parse() {
	// multiTempString := ""
	for {
		if !(p.now < p.length) {
			break
		}
		line := (*p.lines)[p.now]
		// 如果没有处理多行注释并且是单行注释
		if !p.multiLine && strings.HasPrefix(line, "// @") {
			res := typeMatchAndDeal(line[4:])
			p.res += res
		}
		p.now += 1
		// 这里将要加入多行注释的处理
	}
}

// 匹配单行注释并处理
func typeMatchAndDeal(content string) string {
	res := ""
	if strings.HasPrefix(content, "method") {
		res += "#### 请求方法" + content[6:] + "\n\n"
	} else if strings.HasPrefix(content, "router") {
		res += "#### 请求URL" + content[6:] + "\n\n"
	} else if strings.HasPrefix(content, "type") {
		res += "#### 传入格式" + content[4:] + "\n\n"
	}
	return res
}

// 处理多行注释
func (p *Parser) dealMutil() {
	res := ""
	startIndex := p.now
	if !p.multiLine {
		// 说明是从一个完整的多行注释开头开始处理
		// 那先把多行注释的头取出来
		header := (*p.lines)[p.now][4:]
		if strings.HasPrefix(header, "Response") {
			res += "#### Response\n\n"
		} else if strings.HasPrefix(header, "Param") {
			res += "#### 参数格式\n\n"
		}
		startIndex += 1
	}
	leftString, finalIndex, over := reachToEnd(*p.lines, startIndex, p.length)
	p.now = finalIndex + 1
	p.multiLine = over
	res += leftString
	p.res += res
}

func reachToEnd(lines []string, startIndex int, length int) (string, int, bool) {
	res := lines[startIndex]
	over := false
	now := startIndex
	for {
		now += 1
		if strings.HasPrefix(lines[now], "*/") {
			// 以多行注释结尾跳出循环，说明注释处理完毕
			over = true
			break
		} else if !(now < length) {
			// 以切片结束跳出循环，说明注释没有处理完毕
			break
		}
		res += lines[now]
	}
	return res, now, over
}

// 将得到的结果写入文件
func (p *Parser) WriteIntoFile() {
	os.MkdirAll(fe.Dest, 0666)
	destPath := fe.Dest + "/" + p.dest + ".md"
	fmt.Println(destPath)
	// 这里应该只有路径错误吧，没有别的错误吧
	file, err := os.Create(destPath)
	if err != nil {
		log.LogError(log.PATH_ERROR, err.Error())
		// fmt.Println(fe.Yellow, "[PATH ERROR]", err.Error(), destPath, fe.Reset)
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(p.Result())
	writer.Flush()
}
