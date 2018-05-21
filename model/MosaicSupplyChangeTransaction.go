package model

type MosaicSupplyChangeTransaction struct {
	TimeStamp  int      `json:"timeStamp"`
	Signature  string   `json:"signature"`
	Fee        int      `json:"fee"`
	Type       int      `json:"type"`
	Deadline   int      `json:"deadline"`
	Version    uint     `json:"version"`
	Signer     string   `json:"signer"`
	SupplyType int      `json:"supplyType"`
	Delta      int      `json:"delta"`
	MosaicId   MosaicId `json:"mosaicId"`
}
