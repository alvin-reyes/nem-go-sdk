package model

type UnconfirmedTransactionMetaDataPair struct {
	UnconfirmedTransactionMetaData UnconfirmedTransactionMetaData `json:"meta"`
	Transaction                    Transaction                    `json:"transaction"`
}
