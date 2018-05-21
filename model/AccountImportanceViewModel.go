package model

type AccountImportanceViewModel struct {
	Address    string `json:"address"`
	Importance struct {
		IsSet  int    `json:"isSet"`
		Score  string `json:"score"`
		Ev     string `json:"ev"`
		Height string `json:"height"`
	} `json:"importance"`
}
