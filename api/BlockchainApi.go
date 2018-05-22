package api

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/nem-sdk-go/model"
)

type BlockchainApi struct {
	BasePath  string
	ApiClient *ApiClient
}

var chainHeight = "/chain/height"
var chainScore = "/chain/score"
var chainLastBlock = "/chain/last-block"
var blockAtPublic = "/block/at/public"
var localChainBlocksAfter = "/local/chain/blocks-after"

//	local
var localChaniBlocksAfter = "/local/chain/blocks-after"

func NewBlockchainApi(basePath string) *BlockchainApi {
	apiClient := NewApiClient(basePath)
	return &BlockchainApi{basePath, apiClient}
}

func (s BlockchainApi) GetChainHeight() (*model.BlockHeight, error) {
	response, err := s.ApiClient.Get(chainHeight)
	responseModel := &model.BlockHeight{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s BlockchainApi) GetChainScore() (*model.BlockChainScore, error) {
	response, err := s.ApiClient.Get(chainScore)
	responseModel := &model.BlockChainScore{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s BlockchainApi) GetChainLastBlock() (*model.Block, error) {
	response, err := s.ApiClient.Get(chainLastBlock)
	responseModel := &model.Block{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s BlockchainApi) GetChainBlockAtPublic(height model.BlockHeight) (*model.Block, error) {
	jsonInput, _ := json.Marshal(height)
	response, err := s.ApiClient.PostJsonByte(blockAtPublic, "application/json", jsonInput)
	responseModel := &model.Block{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s BlockchainApi) GetLocalChainBlocksAfter(height model.BlockHeight) (*model.ExplorerBlockViewModel, error) {
	jsonInput, _ := json.Marshal(height)
	response, err := s.ApiClient.PostJsonByte(localChainBlocksAfter, "application/json", jsonInput)
	responseModel := &model.ExplorerBlockViewModel{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}
