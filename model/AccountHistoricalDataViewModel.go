package model

type AccountHistoricalDataViewModel struct {
	Height          string `json:"height"`
	Address         string `json:"page"`
	Balance         uint32 `json:"balance"`
	VestedBalance   uint32 `json:"vestedBalance"`
	UnvestedBalance uint32 `json:"unvestedBalance"`
	Importance      uint32 `json:"importance"`
	PageRank        uint32 `json:"pageRank"`
}
