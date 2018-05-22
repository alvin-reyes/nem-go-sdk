package api

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/nem-sdk-go/model"
)

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

func (s NodeApi) GetNodeInfo() (*model.Node, error) {
	response, err := s.ApiClient.Get(nodeInfo)
	responseModel := &model.Node{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s NodeApi) GetNodeExtendedInfo() (*model.NisNodeInfo, error) {
	response, err := s.ApiClient.Get(nodeExtendedInfo)
	responseModel := &model.NisNodeInfo{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s NodeApi) GetNodePeerListAll() (*model.NodeCollection, error) {
	response, err := s.ApiClient.Get(nodePeerListAll)
	responseModel := &model.NodeCollection{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s NodeApi) GetNodePeerListReachable() (*model.NodeCollection, error) {
	response, err := s.ApiClient.Get(nodePeerListReachable)
	responseModel := &model.NodeCollection{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s NodeApi) GetNodePeerListActive() (*model.NodeCollection, error) {
	response, err := s.ApiClient.Get(nodePeerListActive)
	responseModel := &model.NodeCollection{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s NodeApi) GetActivePeersMaxChainHeigh() (*model.BlockHeight, error) {
	response, err := s.ApiClient.Get(nodeActivePeersMaxChainHeight)
	responseModel := &model.BlockHeight{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s NodeApi) GetNodeExperiences() (*model.ExtendedNodeExperiencePair, error) {
	response, err := s.ApiClient.Get(nodeExperiences)
	responseModel := &model.ExtendedNodeExperiencePair{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s NodeApi) GetNodeBoot(bootNodeRequest model.BootNodeRequest) (*model.ExtendedNodeExperiencePair, error) {
	response, err := s.ApiClient.Get(nodeBoot)
	responseModel := &model.ExtendedNodeExperiencePair{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}
