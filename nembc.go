package nembc

import ("github.com/nem-sdk-go/api")

//	this is the client configuration object.
type ClientConfig struct {
	BasePath string 
	Network byte
	AccountApi *api.AccountApi 
	BlockchainApi *api.BlockchainApi
	NamespaceMosaicsApi *api.NamespaceMosaicsApi 
	NodeApi *api.NodeApi
	TransactionApi *api.TransactionApi
}

const (
	Testnet = byte(0x98)
	Mainnet = byte(0x68)
)

//	this creates a new new nis instance which can then be used 
//	to access all services.
func NewNisInstance(basePath string, network byte) *ClientConfig {
	//	create the config. It'll be used for all objects.
	newNisInstance := &ClientConfig {	
		BasePath : basePath,
		Network: network,
		AccountApi: api.NewAccountApi(basePath),
		BlockchainApi: api.NewBlockchainApi(basePath),
		NamespaceMosaicsApi: api.NewNamespaceMosaicsApi(basePath),
		NodeApi: api.NewNodeApi(basePath),
		TransactionApi: api.NewTransactionApi(basePath),
	}
	return newNisInstance
}
