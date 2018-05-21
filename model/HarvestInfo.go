package model

type HarvestInfo struct {
	TimeStamp  string `json:"timeStamp"`
	Id         string `json:"id"`
	Difficulty string `json:"difficulty"`
	TotalFee   string `json:"totalFee"`
	Height     string `json:"height"`
}
