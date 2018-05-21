package api

import (
	"encoding/json"
	"io/ioutil"
	"log"

	model "github.com/nem-sdk-go/model"
)

type AccountApi struct {
	BasePath  string
	ApiClient *ApiClient
}

var accountGenerate = "/account/generate"
var accountGet = "/account/get"
var accountGetFromPublicKey = "/account/get/from-public-key"
var accountGetForwardedFromPublicKey = "/account/get/forwarded/from-public-key"
var accountStatus = "/account/status"

func NewAccountApi(basePath string) *AccountApi {
	apiClient := NewApiClient(basePath)
	return &AccountApi{basePath, apiClient}
}

func (s AccountApi) AccountGenerate() (*model.KeyPairViewModel, error) {
	response, err := s.ApiClient.Get(accountGenerate)
	responseModel := &model.KeyPairViewModel{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err

}

func (s AccountApi) AccountGet(address string) (*model.AccountMetaDataPair, error) {
	response, err := s.ApiClient.Get(accountGet + "?address=" + address)
	responseModel := &model.AccountMetaDataPair{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s AccountApi) AccountGetFromPublicKey(publicKey string) (*model.AccountMetaDataPair, error) {
	response, err := s.ApiClient.Get(accountGetFromPublicKey + "?publicKey=" + publicKey)
	responseModel := &model.AccountMetaDataPair{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s AccountApi) AccountGetForwardedFromPublicKey(publicKey string) (*model.AccountMetaDataPair, error) {
	response, err := s.ApiClient.Get(accountGetForwardedFromPublicKey + "?publicKey=" + publicKey)
	responseModel := &model.AccountMetaDataPair{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s AccountApi) AccountStatus(address string) (*model.AccountMetaData, error) {
	response, err := s.ApiClient.Get(accountStatus + "?address=" + address)
	responseModel := &model.AccountMetaData{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}
