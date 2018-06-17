package model

type MultisigAggregateModificationTransaction struct {
	TimeStamp                       int                                     `json:"timeStamp"`
	Signature                       string                                  `json:"signature"`
	Fee                             int                                     `json:"fee"`
	Type                            int                                     `json:"type"`
	Deadline                        int                                     `json:"deadline"`
	Version                         uint                                    `json:"version"`
	Signer                          string                                  `json:"signer"`
	MultisigCosignatoryModification []MultisigCosignatoryModification `json:"modifications"`
	MinCosignatories                struct {
		RelativeChange string `json:"relativeChange"`
	} `json:"minCosignatories"`
}

type RequestPrepareAnnounceMultisigAggregateModificationTransaction struct {
	MultisigAggregateModificationTransaction MultisigAggregateModificationTransaction `json:"transaction"`
	PrivateKey                               string                                   `json:"privateKey"`
}
