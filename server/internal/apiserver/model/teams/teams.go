package modelTeams

type Team struct {
	Id     int    `json:"id"`
	IdUser string `json:"idUser"`
	Name   string `json:"nameteam"`
	Sport  string `json:"term"`
	Place  string `json:"place"`
}
