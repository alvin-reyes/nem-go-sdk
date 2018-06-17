package model

type ImportanceTransferTransaction struct {
	TimeStamp     string `json:"timeStamp"`
	Signature     string `json:"signature"`
	Fee           int    `json:"fee"`
	Mode          int    `json:"mode"`
	RemoteAccount string `json:"remoteAccount"`
	Type          int    `json:"type"`
	Deadling      int    `json:"deadline"`
	Version       int    `json:"version"`
	Signer        string `json:"signer"`
}

type RequestPrepareAnnounceImportanceTransferTransaction struct {
	ImportanceTransferTransaction ImportanceTransferTransaction `json:"transaction"`
	PrivateKey                    string                        `json:"privateKey"`
}
