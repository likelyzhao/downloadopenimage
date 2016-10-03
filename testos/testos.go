package loadInfos

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type OpenImageInfo struct {
	FileIdx   string
	FileURL   string
	FileTitle string
	Allstring string
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
		splits[len(splits)-1] = SubString(splits[len(splits)-1], 1, len(splits[len(splits)-1])-3)
		infos = append(infos, OpenImageInfo{splits[0], splits[1], splits[len(splits)-1], line})

		defer func() {
			f.Close()
			return
		}()
	}

}

func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	if length < 0 {
		return str
	}
	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth - 1
	}
	// 返回子串
	//fmt.Printf("begin = %d end = %d length = %d\n", begin, end, length)
	return string(rs[begin:end])
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
