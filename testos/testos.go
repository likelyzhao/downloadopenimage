package testos

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//import "os"

func TestfileOpen(filePath string) (flag int) {
	flag = 0
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0660)
	if err != nil {
		flag = 1
		panic(err)
	}
	// stuct inputs
	bfRd := bufio.NewReader(f)
	line, err := bfRd.ReadString('\n')
	if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
		if err == io.EOF {
			return
		}
		return
	}
	splits := strings.Split(line, ",")
	//genSplit
	//line.Split()
	for i := 0; i < len(splits); i++ {
		fmt.Printf("%s\n", splits[i])
	}

	defer f.Close()
	return flag
}

func Add(num1 int, num2 int) (result int) {
	return num1 + num2

}
