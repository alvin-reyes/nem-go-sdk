package model

type NemAnnounceResult struct {
	Type            int    `json:"type"`
	Code            int    `json:"code"`
	Message         string `json:"message"`
	TransactionHash struct {
		Data string `json:"data"`
	} `json:"transactionHash"`
	InnerTransactionHash struct {
		Data string `json:"data"`
	} `json:"innerTransactionHash"`
}
