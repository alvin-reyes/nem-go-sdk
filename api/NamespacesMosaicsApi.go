package api

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/nem-sdk-go/model"
)

type NamespaceMosaicsApi struct {
	BasePath  string
	ApiClient *ApiClient
}

var namespaceRootPage = "/namespace/root/page"
var namespace = "/namespace"
var namespaceMosaicsDefinitionPage = "/namespace/mosaic/definition/page"

func NewNamespaceMosaicsApi(basePath string) *NamespaceMosaicsApi {
	apiClient := NewApiClient(basePath)
	return &NamespaceMosaicsApi{basePath, apiClient}
}

func (s NamespaceMosaicsApi) GetNamespaceRootPage(namespace string) (*model.Namespace, error) {
	response, err := s.ApiClient.Get(namespaceRootPage + "?namespace=" + namespace)
	responseModel := &model.Namespace{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}

func (s NamespaceMosaicsApi) GetNamespaceMosaicDefinitionPage(namespace string) (*model.Namespace, error) {
	response, err := s.ApiClient.Get(namespaceRootPage + "?namespace=" + namespace)
	responseModel := &model.Namespace{}
	data, _ := ioutil.ReadAll(response.Body)
	jsonResult := string(data)
	unmarshalErr := json.Unmarshal([]byte(jsonResult), responseModel)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return responseModel, err
}
