package model

type ExtendedNodeExperiencePair struct {
	Node       Node   `json:"node"`
	Syncs      string `json:"syncs"`
	Experience struct {
		S string `json:"s"`
		F string `json:"f"`
	} `json:"experience"`
}
