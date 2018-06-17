package model

type MultisigTransaction struct {
	TimeStamp  int                            `json:"timeStamp"`
	Signature  string                         `json:"signature"`
	Fee        int                            `json:"fee"`
	Type       int                            `json:"type"`
	Deadline   int                            `json:"deadline"`
	Version    uint                           `json:"version"`
	Signer     string                         `json:"signer"`
	OtherTrans TransferTransactionV2          `json:"otherHash"`
	Signatures []MultisigSignatureTransaction `json:"signatures"`
}

type RequestPrepareAnnounceMultisigTransaction struct {
	MultisigTransaction MultisigTransaction `json:"transaction"`
	PrivateKey          string              `json:"privateKey"`
}
