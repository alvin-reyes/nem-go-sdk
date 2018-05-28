package model

import "github.com/nem-sdk-go/core/transaction"

type UnconfirmedTransactionMetaDataPair struct {
	UnconfirmedTransactionMetaData UnconfirmedTransactionMetaData `json:"meta"`
	Transaction                    transaction.Transaction        `json:"transaction"`
}
