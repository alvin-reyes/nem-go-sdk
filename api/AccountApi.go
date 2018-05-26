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

//	account
var accountGenerate = "/account/generate"
var accountGet = "/account/get"
var accountGetFromPublicKey = "/account/get/from-public-key"
var accountGetForwardedFromPublicKey = "/account/get/forwarded/from-public-key"
var accountStatus = "/account/status"
var accountIncomingTransaction = "/account/transfers/incoming"
var accountOutgoingTransaction = "/account/transfers/outgoing"
var accountAllTransaction = "/account/transfers/all"
var accountUnconfirmedTransaction = "/account/unconfirmedTransactions"

//	local
var localAccountIncomingTransaction = "/local/account/transfers/incoming"
var localAccountOutgoingTransaction = "/local/account/transfers/outgoing"
var localAccountAllTransaction = "/local/account/transfers/all"

// harvets and importances
var accountHarvest = "/account/harvests"
var accountImportances = "/account/importances"
var accountNamespacePage = "/account/namespace/page"
var accountMosaicDefinitionPage = "/account/mosaic/definition/page"
var accountMosaicOwned = "/account/mosaic/owned"

//	lock
var accountUnlock = "/account/unlock"
var accountLock = "/account/lock"
var accountUnlockInfo = "/account/unlocked/info"

//	historical
var accountHistoricalGet = "/account/historical/get"

func NewAccountApi(basePath string) *AccountApi {
	apiClient := NewApiClient(basePath)
	return &AccountApi{basePath, apiClient}
}

func (s AccountApi) AccountGenerate() (*model.KeyPairViewModel, error) {
	response, err := s.ApiClient.Get(accountGenerate)

	if err != nil 
		return 
		
	responseModel := &model.KeyPairViewModel{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		return
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

func (s AccountApi) GetAccountIncomingTransactionsUsingAddress(address string) (*model.TransactionMetaDataPair, error) {
	addressOnly := accountIncomingTransaction + "?address=" + address
	return s.getTransaction(addressOnly)
}

func (s AccountApi) GetAccountIncomingTransactionsUsingAddressHash(address string, hash string) (*model.TransactionMetaDataPair, error) {
	addressOnly := accountIncomingTransaction + "?address=" + address + "&hash=" + hash
	return s.getTransaction(addressOnly)
}

func (s AccountApi) GetAccountIncomingTransactionsUsingAdderssHashId(address string, hash string, id string) (*model.TransactionMetaDataPair, error) {
	addressOnly := accountIncomingTransaction + "?address=" + address + "&hash=" + hash + "&id=" + id
	return s.getTransaction(addressOnly)
}

func (s AccountApi) GetAccountOutgoingTransactionsUsingAddress(address string) (*model.TransactionMetaDataPair, error) {
	addressOnly := accountOutgoingTransaction + "?address=" + address
	return s.getTransaction(addressOnly)
}

func (s AccountApi) GetAccountOutgoingTransactionsUsingAddressHash(address string, hash string) (*model.TransactionMetaDataPair, error) {
	addressOnly := accountOutgoingTransaction + "?address=" + address + "&hash=" + hash
	return s.getTransaction(addressOnly)
}

func (s AccountApi) GetAccountOutgoingTransactionsUsingAdderssHashId(address string, hash string, id string) (*model.TransactionMetaDataPair, error) {
	addressOnly := accountOutgoingTransaction + "?address=" + address + "&hash=" + hash + "&id=" + id
	return s.getTransaction(addressOnly)
}

func (s AccountApi) GetAccountAllTransactionsUsingAddress(address string) (*model.TransactionMetaDataPair, error) {
	addressOnly := accountAllTransaction + "?address=" + address
	return s.getTransaction(addressOnly)
}

func (s AccountApi) GetAccountAllTransactionsUsingAddressHash(address string, hash string) (*model.TransactionMetaDataPair, error) {
	addressOnly := accountAllTransaction + "?address=" + address + "&hash=" + hash
	return s.getTransaction(addressOnly)
}

func (s AccountApi) GetAccountAllTransactionsUsingAdderssHashId(address string, hash string, id string) (*model.TransactionMetaDataPair, error) {
	addressOnly := accountAllTransaction + "?address=" + address + "&hash=" + hash + "&id=" + id
	return s.getTransaction(addressOnly)
}

func (s AccountApi) GetAccountUnconfirmedTransactions(address string) (*model.TransactionMetaDataPair, error) {
	addressOnly := accountUnconfirmedTransaction + "?address=" + address
	return s.getTransaction(addressOnly)
}

func (s AccountApi) getTransaction(endpoint string) (*model.TransactionMetaDataPair, error) {
	response, err := s.ApiClient.Get(endpoint)
	responseModel := &model.TransactionMetaDataPair{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}
