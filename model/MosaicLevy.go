package model

type MosaicLevy struct {
	Type      int      `json:"type"`
	Recipient string   `json:"recipient"`
	MosaicId  MosaicId `json:"mosaicId"`
	Fee       int      `json:"fee"`
}
