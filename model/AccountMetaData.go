package model

type AccountMetaData struct {
	Status        string        `json:"status"`
	RemoteStatus  string        `json:"remoteStatus"`
	CosignatoryOf []AccountInfo `json:"cosignatoryOf"`
	Cosignatories []AccountInfo `json:"cosignatories"`
}
