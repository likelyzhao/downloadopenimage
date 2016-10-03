package download

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"testos"
)

var urlist = [...]string{"http://stock.591hx.com/article/2014-12-03/0000850005s.shtml"}
var album chan string
var w sync.WaitGroup
var logger *log.Logger
var dir string

func getImageFromURL(url string, savepath string) error {
	data := GetUrl(url)
	album <- "OK"
	//	os.exis(savepath)
	f, err := os.Create(savepath)
	if err != nil {
		if true == os.IsExist(err) {
			return err
		}
		panic(err)
	}
	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}

	f.Close()
	defer func() {
		f.Close()
		w.Done()
	}()
	return err
}

func TestMain(infos []loadInfos.OpenImageInfo, savedir string) {
	//dir = "tmp_chenjo/"
	err := os.Mkdir(savedir, 0777)
	if err != nil {
		isexist := os.IsExist(err)
		log.Println(isexist)
	}
	f, err := os.Create("log.txt")
	logger = log.New(f, "", 0)

	//	for i := 0; i < 10; i++ {

	//	}

	album = make(chan string, 200)
	for _, v := range infos {
		saveimagepath := savedir + v.FileIdx + ".jpg"
		w.Add(1)
		go getImageFromURL(v.FileURL, saveimagepath)
		logger.Println(v.Allstring, <-album)
		//w.Wait()
	}

	f.Close()

	//for _, v := range infos {
	//	w.Add(1)
	//	go GetAlbum(v)
	//	w.Wait()
	//}
}

func GetAlbum(url string) {
	data := GetUrl(url)
	body := string(data)
	//

	part := regexp.MustCompile(``)
	match := part.FindAllStringSubmatch(body, -1)
	for _, v := range match {

		if m, _ := regexp.MatchString(`.*/hnimg/201412/03/.*\.jpg`, v[1]); !m {
			continue
		}
		//println(v[1])
		album <- v[1]
		w.Add(1)
		go GetItem()
	}
	w.Done()

}

func GetItem() {
	url := <-album
	println(url)
	defer func() {
		ret := recover()
		if ret != nil {
			log.Println(ret)
			w.Done()
		} else {
			w.Done()
		}
	}()

	//data := GetUrl(url)
	//if len(data) > 10 {
	//body := string(data)
	//part := regexp.MustCompile(`bigimgsrc="(.*)"`)
	//match := part.FindAllStringSubmatch(body, -1)
	//for _, v := range match {
	str := strings.Split(url, "/")
	length := len(str)
	source := GetUrl(url)
	name := str[length-1]
	file, err := os.Create(dir + name)
	if err != nil {
		panic(err)
	}
	size, err := file.Write(source)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	log.Println(size)
	//}
	//}
}

func GetUrl(url string) []byte {
	ret, err := http.Get(url)
	if err != nil {
		log.Println(url)
		status := map[string]string{}
		status["status"] = "400"
		status["url"] = url
		panic(status)
	}
	body := ret.Body
	fmt.Printf("begin reading %s\n", url)
	data, _ := ioutil.ReadAll(body)
	fmt.Printf("end reading %s\n", url)
	return data
}
