package main

import ServerClient "Smart-city/third_party/client/internal"

func main() {
	URL := ServerClient.DownloadNews()
	if URL != nil {
		for _, url := range URL {
			ServerClient.GetPic(url)
		}
	}
}
