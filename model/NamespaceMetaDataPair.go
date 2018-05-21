package model

type NamespaceMetaDataPair struct {
	Meta struct {
		Id string `json:"id"`
	} `json:"meta"`
	Namespace Namespace `json:"namespace"`
}
