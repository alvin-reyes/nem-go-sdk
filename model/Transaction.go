package model

const (
	TransferTransactionType                              = 0x0101
	ImportantTransactionTransactionType                  = 0x0801
	MultisigAggregateModificationTransferTransactionType = 0x1001
	MultisigSignatureTransactionType                     = 0x1002
	MultisigTransactionType                              = 0x1004
	ProvisionNamespaceTransactonType                     = 0x2001
	MosaicDefinitionCreationTransactionType              = 0x4001
	MosaicSupplyChangeTransactionType                    = 0x4002
)

type Transaction struct {
	TimeStamp string `json:"timeStamp"`
	Signature string `json:"signature"`
	Fee       string `json:"fee"`
	Type      string `json:"type"`
	Deadlne   string `json:"deadline"`
	Version   string `json:"version"`
	Signer    string `json:"signer"`
}
