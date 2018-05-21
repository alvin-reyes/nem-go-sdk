package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ApiClient struct {
	basePath string
}

func NewApiClient(basePath string) *ApiClient {
	return &ApiClient{basePath}
}

func (s ApiClient) Get(endpoint string) (*http.Response, error) {
	return http.Get(s.basePath + endpoint)
}

func (s ApiClient) Post(endpoint string, contentType string, jsonMap map[string]string) (*http.Response, error) {
	jsonValue, _ := json.Marshal(jsonMap)
	return http.Post(s.basePath+endpoint, contentType, bytes.NewBuffer(jsonValue))
}

// func callWS(name string, lastHash string, pvkey string, pbkey string, messageType string) string {
// 	fmt.Println("Starting the application...")
// 	jsonData := map[string]string{
// 		"contentType": "string",
// 		"ipfsHash":    lastHash,
// 		"keywords":    "string",
// 		"messageType": messageType,
// 		"name":        name,
// 		"privateKey":  pvkey,
// 		"publicKey":   pbkey,
// 	}

// 	jsonValue, _ := json.Marshal(jsonData)
// 	response, err := http.Post("http://localhost:8881/upload/local/resource", "application/json", bytes.NewBuffer(jsonValue))

// 	if err != nil {
// 		fmt.Printf("The HTTP request failed with error %s\n", err)
// 		return ""
// 	} else {
// 		data, _ := ioutil.ReadAll(response.Body)
// 		fmt.Println(string(data))
// 		return string(data)
// 	}
// }
