package modelNews

type News struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Txt    string `json:"text"`
	PicURL string `json:"picture"`
	Time   int64  `json:"time"`
}
