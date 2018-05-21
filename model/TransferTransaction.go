package model

type TransferTransaction struct {
	TimeStamp string
	Signature string
	Fee       string
	Type      string
	Deadlne   string
	Version   string
	Signer    string

	Recipient   string
	Message     string
	Payload     string
	MessageType string
}
