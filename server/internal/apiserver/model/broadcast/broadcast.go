package modelBroadcast

type Broadcast struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	BroadURL string `json:"broadURL"`
	PicURL   string `json:"picture"`
}
