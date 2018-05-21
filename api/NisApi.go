package api

import (
	"fmt"
	"io/ioutil"
)

var nisHeartBeat = "/heartbeat"
var nisStatus = "/status"

type NisApi struct {
	BasePath  string
	ApiClient *ApiClient
}

func NewNisApi(basePath string) *NisApi {
	apiClient := NewApiClient(basePath)
	return &NisApi{basePath, apiClient}
}

func (s NisApi) GetHeartBeat() string {
	response, err := s.ApiClient.Get(nisHeartBeat)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}

}

func (s NisApi) GetStatus() string {
	response, err := s.ApiClient.Get(nisStatus)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}
}
