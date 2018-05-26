package model

import (
	"github.com/nem-sdk-go/model"
)

type Block struct {
	Timestamp     string `json:"timeStamp"`
	Signature     string `json:"signature"`
	PrevBlockHash struct {
		Data string `json:"data"`
	} `json:"prevBlockHash"`
	Type         string              `json:"type"`
	Transactions []model.Transaction `json:"transactions"`
	Version      string              `json:"version"`
	Signer       string              `json:"signer"`
	Height       string              `json:"height"`
}
