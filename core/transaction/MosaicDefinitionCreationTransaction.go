package transaction

import "github.com/nem-sdk-go/model"

type MosaicDefinitionCreationTransaction struct {
	TimeStamp        int                    `json:"timeStamp"`
	Signature        string                 `json:"signature"`
	Fee              int                    `json:"fee"`
	Type             int                    `json:"type"`
	Deadline         int                    `json:"deadline"`
	version          uint                   `json:"version"`
	Signer           string                 `json:"signer"`
	CreationFee      int                    `json:"creationFee"`
	CreationFeeSink  string                 `json:"creationFeeSink"`
	MosaicDefinition model.MosaicDefinition `json:"mosaicDefinition"`
	Levy             model.Levy             `json:"levy"`
}

type RequestPrepareAnnounceMosaicDefinitionCreationTransaction struct {
	MosaicDefinitionCreationTransaction MosaicDefinitionCreationTransaction `json:"transaction"`
	PrivateKey                          string                              `json:"privateKey"`
}
