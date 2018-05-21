package model

type TransferTransaction struct {
	TimeStamp string `json:"timeStamp"`
	Signature string `json:"signature"`
	Fee       string `json:"fee"`
	Type      string `json:"type"`
	Deadlne   string `json:"deadline"`
	Version   string `json:"version"`
	Signer    string `json:"signer"`

	Recipient   string `json:"recipient"`
	Message     string `json:"message"`
	Payload     string `json:"payload"`
	MessageType string `json:"messageType"`
}
