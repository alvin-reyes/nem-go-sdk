package model

type ExplorerBlockViewModel struct {
	Data ExplorerBlockViewModelData `json:"data"`
}

type ExplorerBlockViewModelData struct {
	Txes      []ExplorerTransferViewModel `json:"txes"`
	Block     string                      `json:"block"`
	Hash      string                      `json:"hash"`
	Difficult string                      `json:"difficulty"`
}
