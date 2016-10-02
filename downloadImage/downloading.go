package download

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

var urlist = [...]string{"http://stock.591hx.com/article/2014-12-03/0000850005s.shtml"}
var album chan string
var w sync.WaitGroup
var dir string

func getImageFromUrl(infos []OpenImageInfo) {

}

func TestMain() {
	dir = "tmp_chenjo/"
	err := os.Mkdir(dir, 0777)
	if err != nil {
		isexist := os.IsExist(err)
		log.Println(isexist)
	}
	album = make(chan string, 200)

	for _, v := range urlist {
		w.Add(1)
		go GetAlbum(v)
		w.Wait()
	}
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
	data, _ := ioutil.ReadAll(body)
	return data
}
