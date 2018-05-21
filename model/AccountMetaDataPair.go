package model

// type AccountMetaDataPair struct {
// 	Status        string
// 	RemoteStatus  string
// 	CosignatoryOf string
// }

type AccountMetaDataPair struct {
	AccountInfo     AccountInfo     `json:"account"`
	AccountMetaData AccountMetaData `json:"meta"`
}
