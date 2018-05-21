package api

type TransactionApi struct {
	BasePath  string
	ApiClient *ApiClient
}

var transactionPrepareAnnounce = "/transaction/prepare-announce"

func NewTransactionApi(basePath string) *TransactionApi {
	apiClient := NewApiClient(basePath)
	return &TransactionApi{basePath, apiClient}
}
