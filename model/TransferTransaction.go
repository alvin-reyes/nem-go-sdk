package model

//	V2
type TransferTransaction struct {
	TimeStamp string `json:"timeStamp"`
	Amount    int    `json:"amount"`
	Signature string `json:"signature"`
	Fee       string `json:"fee"`
	Recipient string `json:"recipient"`
	Type      int    `json:"type"`
	Deadline  string `json:"deadline"`
	Message   struct {
		Payload string `json:"payload"`
		Type    int    `json:"type"`
	} `json:"message"`
	Version string `json:"version"`
	Signer  string `json:"signer"`

	Mosaics []MosaicIdQuantity `json:"mosaics"`
}

type MosaicIdQuantity struct {
	MosaicId MosaicId `json:"mosaicId"`
	Quantity int      `json:"quantity"`
}
