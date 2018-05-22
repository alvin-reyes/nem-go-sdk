package model

type MosaicDefinition struct {
	Creator string `json:"creator"`
	Id      struct {
		NamespaceId string `json:"namespaceId"`
	} `json:"id"`
}
