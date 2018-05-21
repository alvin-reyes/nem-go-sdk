package api

type BlockchainApi struct {
	BasePath  string
	ApiClient *ApiClient
}

var chainHeight = "/chain/height"
var chainScore = "/chain/score"
var chainLastBlock = "/chain/last-block"
var blockAtPublic = "/block/at/public"

//	local
var localChaniBlocksAfter = "/local/chain/blocks-after"

func NewBlockchainApi(basePath string) *BlockchainApi {
	apiClient := NewApiClient(basePath)
	return &BlockchainApi{basePath, apiClient}
}
