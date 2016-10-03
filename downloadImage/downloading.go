package download

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"testos"
	"time"
)

var album chan string
var w sync.WaitGroup
var logger *log.Logger

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getImageFromURL(url string, savepath string) error {
	fmt.Printf("begin reading %s\n", url)
	flag, _ := PathExists(savepath)
	if flag == true {
		return nil
	}

	data, err := getUrl(url)
	if err != nil {
		album <- "error"
		panic(err)
	}
	//data := []byte{1, 2}
	album <- "OK"
	//	os.exis(savepath)

	f, err := os.Create(savepath)
	if err != nil {
		if true == os.IsExist(err) {
			album <- "error"
			return err
		}
		panic(err)
	}
	_, err = f.Write(data)
	if err != nil {
		album <- "error"
		panic(err)
	}

	defer func() {
		fmt.Println(recover())
		defer f.Close()
		log.Println(album)
		w.Done()
	}()
	fmt.Printf("end reading %s\n", url)
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
	album = make(chan string, len(infos))
	for _, v := range infos {
		saveimagepath := savedir + v.FileIdx + ".jpg"
		w.Add(1)
		go getImageFromURL(v.FileURL, saveimagepath)
		time.Sleep(time.Second / 10)
		//logger.Println(v.Allstring + <-album)
		//w.Wait()
	}
	w.Wait()
	for i := 0; i < len(infos); i++ {
		logger.Println(infos[i].Allstring + <-album)
	}

	f.Close()

	//for _, v := range infos {
	//	w.Add(1)
	//	go GetAlbum(v)
	//	w.Wait()
	//}
}

func getUrl(url string) (data []byte, err error) {
	ret, err := http.Get(url)
	if err != nil {
		log.Println(url)
		return data, err
	}
	body := ret.Body

	data, _ = ioutil.ReadAll(body)

	return data, nil
}
