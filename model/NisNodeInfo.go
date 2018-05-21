package model

type NisNodeInfo struct {
	Node    Node                `json:"node"`
	NisInfo ApplicationMetaData `json:"nisInfo"`
}
