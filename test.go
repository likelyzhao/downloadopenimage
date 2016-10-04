package main

import (
	download "downloadImage"
	"fmt"
	"testos"
)

func main() {
	//testos()
	//var res
	//var filename = "E:\\Research\\open Image\\images_2016_08_v2\\train\\images.csv"
	var filename = "part0.txt"
	//res := testos.TestfileOpen(filename)
	res := loadInfos.LoadingOpenImageInfo(filename)

	////spilts the total lists
	/// total numer is 9M
	/// split into 9 parts
	/*
		var flist []*os.File
		for i := 0; i < 10; i++ {
			partFileName := fmt.Sprintf("part%d.txt", i)
			f, _ := os.Create(partFileName)
			flist = append(flist, f)
		}

		idx := 0
		tempsum := 0
		for _, v := range res {
			flist[idx].WriteString(v.Allstring)
			tempsum++
			if tempsum >= 1000000 {
				tempsum = 0
				idx++
			}
		}

		for i := 0; i < 10; i++ {
			flist[i].Close()
		}
	*/
	download.TestMain(res, "E:\\down2\\")
	//	download.TestMain()

	fmt.Printf("openfile_flag= %d\n", len(res))
	fmt.Println("test")
}
