package ServerClient

import (
	modelBroadcast "Smart-city/internal/apiserver/model/broadcast"
	modelNews "Smart-city/internal/apiserver/model/news"
	modeltimetable "Smart-city/internal/apiserver/model/timetable"
	"bytes"
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

func GetPic(url string, urlfull string) {
	res, err := http.Get(strings.Join([]string{"http://localhost:8080", urlfull}, "/"))

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

func GetBroadcast() []string {
	resp, err := http.Get("http://localhost:8080/broadcast")
	if err != nil {
		log.Fatalf("Got error %s", err.Error())
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	results := []modelBroadcast.Broadcast{}
	json.Unmarshal(body, &results)

	URL := []string{}

	for _, result := range results {
		log.Println(result)
		URL = append(URL, result.PicURL)
	}

	return URL
}

func UploadTimetable(idUser int) {
	message := &modeltimetable.Timetable{
		IdUser: idUser,
		Title:  "title",
		Txt:    "txt",
		Time:   1654009999,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
		return
	}

	resp, err := http.Post("http://localhost:8080/upload/timetable", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	var result int
	json.Unmarshal(body, &result)

	log.Println(result)
}

func GetTimetable(idUser int) {
	message := &modeltimetable.Timetable{
		IdUser: idUser,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
		return
	}

	resp, err := http.Post("http://localhost:8080/timetabel", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	results := []modeltimetable.Timetable{}
	json.Unmarshal(body, &results)

	for _, result := range results {
		log.Println(result)
	}
}
