package transaction

type ProvisionNamespaceTransaction struct {
	TimeStamp     int    `json:"timeStamp"`
	Signature     string `json:"signature"`
	Fee           int    `json:"fee"`
	Type          int    `json:"type"`
	Deadline      int    `json:"deadline"`
	Version       uint   `json:"version"`
	Signer        string `json:"signer"`
	RentalFeeSink string `json:"rentalFeeSink"`
	RentalFee     int    `json:"rentalFee"`
	NewPart       string `json:"newPart"`
	Parent        string `json:"parent"`
}
