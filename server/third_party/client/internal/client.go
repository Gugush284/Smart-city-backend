package ServerClient

import (
	modelNews "Smart-city/internal/apiserver/model/news"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func DownloadNews() []string {

	resp, err := http.Get("http://localhost:8080/news")
	if err != nil {
		log.Fatalf("Got error %s", err.Error())
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	results := []modelNews.News{}

	json.Unmarshal(body, &results)

	URL := []string{}

	for _, result := range results {
		log.Println(result.Time)
		URL = append(URL, result.PicURL)
	}

	return URL
}

func GetPic(url string) {
	res, err := http.Get(strings.Join([]string{"http://localhost:8080/news", url}, "/"))

	if err != nil {
		log.Fatalf("Got error %s", err.Error())
		return
	}

	//log.Println(res)
	//log.Println(res.Body)

	defer res.Body.Close()

	path := strings.Join([]string{"third_party/client/download", url}, "/")
	img, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(img, res.Body)
	if err != nil {
		log.Fatal(err)
	}
}
