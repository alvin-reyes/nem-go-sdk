package api

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/nem-sdk-go/model"
)

type TransactionApi struct {
	BasePath  string
	ApiClient *ApiClient
}

var transactionPrepareAnnounce = "/transaction/prepare-announce"
var transactionAnnounce = "/transaction/announce"

func NewTransactionApi(basePath string) *TransactionApi {
	apiClient := NewApiClient(basePath)
	return &TransactionApi{basePath, apiClient}
}

func (s TransactionApi) TransactionPrepareAnnounce(requestPrepareAnnounce model.RequestPrepareAnnounce) (*model.NemAnnounceResult, error) {
	jsonInput, _ := json.Marshal(requestPrepareAnnounce)
	response, err := s.ApiClient.PostJsonByte(transactionPrepareAnnounce, "application/json", jsonInput)
	responseModel := &model.NemAnnounceResult{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s TransactionApi) TransactionAnnounce(requestAnnounce model.RequestAnnounce) (*model.RequestAnnounce, error) {
	jsonInput, _ := json.Marshal(requestAnnounce)
	response, err := s.ApiClient.PostJsonByte(transactionAnnounce, "application/json", jsonInput)
	responseModel := &model.RequestAnnounce{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}
