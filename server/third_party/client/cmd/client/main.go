package main

import (
	ServerClient "Smart-city/third_party/client/internal"
)

func main() {
	/*URL := ServerClient.DownloadNews()
	for _, url := range URL {
		ServerClient.GetPic(url, strings.Join([]string{"news", url}, "/"))
	}

	URL = ServerClient.GetBroadcast()
	for _, url := range URL {
		ServerClient.GetPic(url, strings.Join([]string{"broadcast", url}, "/"))
	}*/

	//id_user := 1
	//ServerClient.UploadTimetable(id_user)
	//ServerClient.GetTimetable(id_user)
	//ServerClient.GetTeams(id_user)
	ServerClient.RegEvent()
}
