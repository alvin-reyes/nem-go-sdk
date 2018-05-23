package model

type RequestPrepareAnnounce struct {
	Transaction struct {
		TimeStamp int    `json:"timeStamp"`
		Amount    int    `json:"amount"`
		Fee       int    `json:"fee"`
		Recipient string `json:"recipient"`
		Type      int    `json:"type"`
		Deadline  int    `json:"deadline"`
		Message   struct {
			Payload string `json:"payload"`
			Type    int    `json:"type"`
		} `json:"message"`
		Version int64  `json:"version"`
		Signer  string `json:"signer"`
	} `json:"transaction"`
	PrivateKey string `json:"privateKey"`
}
