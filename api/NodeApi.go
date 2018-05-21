package api

type NodeApi struct {
	BasePath  string
	ApiClient *ApiClient
}

var nodeInfo = "/node/info"
var nodeExtendedInfo = "/node/extended-info"
var nodePeerListAll = "/node/peer-list/all"
var nodePeerListReachable = "/node/peer-list/reachable"
var nodePeerListActive = "/node/peer-list/active"
var nodeActivePeersMaxChainHeight = "/node/active-peers/max-chain-height"
var nodeExperiences = "/node/experiences"
var nodeBoot = "/node/boot"

func NewNodeApi(basePath string) *NodeApi {
	apiClient := NewApiClient(basePath)
	return &NodeApi{basePath, apiClient}
}
