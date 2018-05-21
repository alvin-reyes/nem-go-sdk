package api

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
