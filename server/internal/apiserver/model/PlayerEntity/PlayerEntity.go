package modelPlayerentity

type Playerentity struct {
	Id         int    `json:"id"`
	Lastname   string `json:"Lastname"`
	Firstname  string `json:"Firstname"`
	Middlename string `json:"Middlename"`
	Age        int    `json:"Age"`
}
