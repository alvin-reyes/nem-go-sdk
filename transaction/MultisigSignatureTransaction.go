package transaction

type MultisigSignatureTransaction struct {
	TimeStamp int    `json:"timeStamp"`
	Signature string `json:"signature"`
	Fee       int    `json:"fee"`
	Type      int    `json:"type"`
	Deadline  int    `json:"deadline"`
	Version   uint   `json:"version"`
	Signer    string `json:"signer"`
	OtherHash struct {
		Data string `json:"data"`
	} `json:"otherHash"`
	OtherAccount string `json:"otherAccount"`
}
