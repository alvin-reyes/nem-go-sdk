package model

type TransactionMetaDataPair struct {
	Meta        TransactionMetaData `json:"meta"`
	Transaction Transaction         `json:"transaction"`
}
