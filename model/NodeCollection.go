package model

type NodeCollection struct {
	Inactive []Node `json:"inactive"`
	Active   []Node `json:"active"`
	Busy     []Node `json:"busy"`
	Failure  []Node `json:"failure"`
	Data     string `json:"data"`
}
