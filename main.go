package main

import (
	"bufio"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mylxsw/go-utils/diff"
	"github.com/mylxsw/go-utils/file"
)

var dataDir, name, inputFile string
var contextLine int

func main() {
	flag.StringVar(&dataDir, "data-dir", "/tmp", "状态数据存储目录")
	flag.StringVar(&name, "name", "default", "当前对比项目名称")
	flag.IntVar(&contextLine, "context-line", 0, "输出上下文行数")
	flag.StringVar(&inputFile, "file", "", "对比的文件名，留空则从标准输入读取")
	flag.Parse()

	var target string
	if inputFile != "" {
		targetBytes, err := ioutil.ReadFile(inputFile)
		if err != nil {
			panic(err)
		}

		target = string(targetBytes)
	} else {
		target = readStdin(100000)
	}

	fs := file.LocalFS{}
	if err := fs.MkDir(dataDir); err != nil {
		panic(err)
	}

	differ := diff.NewDiffer(fs, dataDir, contextLine)
	res := differ.DiffLatest(name, target)
	if err := res.PrintAndSave(os.Stdout); err != nil {
		panic(err)
	}
}

func readStdin(maxLines int) string {
	result := ""

	reader := bufio.NewReader(os.Stdin)
	lineNo := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		lineNo++
		if lineNo > maxLines {
			break
		}

		result += line
	}

	return strings.TrimSpace(strings.TrimSuffix(result, "\n"))
}
