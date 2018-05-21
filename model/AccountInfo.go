package model

type AccountInfo struct {
	Address         string  `json:"address"`
	Balance         int     `json:"balance"`
	VestedBalance   int     `json:"vestedBalance"`
	Importance      float32 `json:"importance"`
	PublicKey       string  `json:"publicKey"`
	Label           string  `json:"label"`
	HarvestedBlocks int     `json:"harvestedBlocks"`
}
