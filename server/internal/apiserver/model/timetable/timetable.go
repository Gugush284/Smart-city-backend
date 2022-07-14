package modeltimetable

type Timetable struct {
	Id     int    `json:"id"`
	IdUser int    `json:"id_user"`
	Title  string `json:"title"`
	Txt    string `json:"text"`
	Time   int    `json:"time"`
}
