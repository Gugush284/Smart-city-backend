package modelScoreboard

type Scoreboard struct {
	FirstTeam    string `json:"firstteam"`
	SecondTeam   string `json:"secondteam"`
	FirstNumber  string `json:"firstNumber"`
	SecondNumber string `json:"secondNumber"`
	Type         string `json:"vid"`
	URL          string `json:"url"`
	Term         string `json:"term"`
}
