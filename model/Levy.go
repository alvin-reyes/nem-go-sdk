package model

type Levy struct {
	Type      string   `json:"type"`
	Recipient string   `json:"recipient"`
	MosaicId  MosaicId `json:"mosaicId"`
	Fee       string   `json:"fee"`
}
