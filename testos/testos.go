package loadInfos

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type OpenImageInfo struct {
	fileIdx   string
	fileUrl   string
	fileTitle string
	stringAll string
}

//import "os"
func LoadingOpenImageInfo(filePath string) (infos []OpenImageInfo) {
	//infosS := make([]OpenImageInfo, 1, 10)
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0660)
	if err != nil {
		panic(err)
	}
	// stuct inputs
	bfRd := bufio.NewReader(f)
	line, err := bfRd.ReadString('\n')
	if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
		if err == io.EOF {
			panic(err)
		}
		panic(err)
	}
	splits := strings.Split(line, ",")
	for i := 0; i < len(splits); i++ {
		fmt.Printf("%s\n", splits[i])
	}

	for {
		line, err := bfRd.ReadString('\n')
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				f.Close()
				return
			}
			panic(err)
		}
		splits := strings.Split(line, ",")
		infos = append(infos, OpenImageInfo{splits[0], splits[1], splits[len(splits)-1], line})

		defer func() {
			f.Close()
			return
		}()
	}

}

/*
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

	for {

		line, err := bfRd.ReadString('\n')
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return
			}
			return
		}
	}

	defer f.Close()
	return flag
}
*/
func Add(num1 int, num2 int) (result int) {
	return num1 + num2

}
