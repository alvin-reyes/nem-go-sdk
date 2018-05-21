package model

type Error struct {
	TimeStamp int    `json:"timeStamp"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Status    string `json:"status"`
}
