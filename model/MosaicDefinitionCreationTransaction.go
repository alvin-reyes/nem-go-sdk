package model

type MosaicDefinitionCreationTransaction struct {
	TimeStamp        int              `json:"timeStamp"`
	Signature        string           `json:"signature"`
	Fee              int              `json:"fee"`
	Type             int              `json:"type"`
	Deadline         int              `json:"deadline"`
	version          uint             `json:"version"`
	Signer           string           `json:"signer"`
	CreationFee      int              `json:"creationFee"`
	CreationFeeSink  string           `json:"creationFeeSink"`
	MosaicDefinition MosaicDefinition `json:"mosaicDefinition"`
}
