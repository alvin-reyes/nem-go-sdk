package model

type TransactionMetaData struct {
	Height string `json:"height"`
	Id     int    `json:"id"`
	Hash   Hash   `json:"hash"`
}
