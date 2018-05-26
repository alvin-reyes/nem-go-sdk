package transaction

import "github.com/nem-sdk-go/model"

type MosaicSupplyChangeTransaction struct {
	TimeStamp  int            `json:"timeStamp"`
	Signature  string         `json:"signature"`
	Fee        int            `json:"fee"`
	Type       int            `json:"type"`
	Deadline   int            `json:"deadline"`
	Version    uint           `json:"version"`
	Signer     string         `json:"signer"`
	SupplyType int            `json:"supplyType"`
	Delta      int            `json:"delta"`
	MosaicId   model.MosaicId `json:"mosaicId"`
}

type RequestPrepareAnnounceMosaicSupplyChangeTransaction struct {
	MosaicSupplyChangeTransaction MosaicSupplyChangeTransaction `json:"transaction"`
	PrivateKey                    string                        `json:"privateKey"`
}
