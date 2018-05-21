package model

type ExplorerTransferViewModel struct {
	Transaction Transaction `json:"tx"`
	Hash        string      `json:"hash"`
	InnerHash   string      `json:"innerHash"`
}
