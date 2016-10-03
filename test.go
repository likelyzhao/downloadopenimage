package main

import (
	download "downloadImage"
	"fmt"
	"testos"
)

func main() {
	//testos()
	//var res
	var filename = "E:\\Research\\open Image\\images_2016_08_v2\\validation\\images.csv"
	//res := testos.TestfileOpen(filename)
	res := loadInfos.LoadingOpenImageInfo(filename)
	download.TestMain(res, "E:\\down\\")
	//	download.TestMain()

	fmt.Printf("openfile_flag= %d\n", len(res))
	fmt.Println("test")
}
