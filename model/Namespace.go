package model

type Namespace struct {
	Fqn    string `json:"fqn"`
	Owner  string `json:"owner"`
	Height string `json:"height"`
}
