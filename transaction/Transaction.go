package transaction

type Transaction struct {
	TimeStamp string `json:"timeStamp"`
	Signature string `json:"signature"`
	Fee       string `json:"fee"`
	Type      string `json:"type"`
	Deadlne   string `json:"deadline"`
	Version   string `json:"version"`
	Signer    string `json:"signer"`
}
