# NEM/NIS Go Wrapper - Simple NEM/NIS1 Wrapper written in Go.

A Simple NEM NIS1 Wrapper written in Go. Basic functions are available such as:

Features:

+ Generation of Account and Lookup
+ Transaction Lookup
+ Blockchain Lookup
+ Node Information Lookup
+ Namespace and Mosaics Lookup
+ Address Generator (with Vanity check)

To follow:

+ CLI Capability
+ Initiated and Sign offline Transactions
+ Initiate and Announce transactions
+ Websocket Monitoring

## How to use

```go
package main

import (
	"github.com/nem-sdk-go/nembc"
)

func main() {

	//	create the instance.
	newNisConnection := nembc.NewNisConnection("http://104.128.226.60:7890/",nembc.Testnet)
	
	//	Get account information
	newNisConnection.AccountApi.AccountGet("TC5KTUA4TTSSCSFQX2DOMCI2BUEHO2CSE4ZCITPM")

	//	Get namespace and mosaic information
    newNisConnection.NamespaceMosaicsApi.GetNamespaceRootPage("roopage")
    
    //  Blockchain Api endpoint
    newNisConnection.BlockchainApi.GetChainHeight()

    //  NisApi Heart Beat
    newNisConnection.NisApi.GetHeartBeat()

    //  Node Information
    newNisConnection.NodeApi.GetNodeInfo()

    // Transaction Api
    newNisConnection.TransactionApi.TransactionPrepareAnnounce()

}
```

## Contribution
The library/tool is far from finish. We need a few contributors to finish some of the tasks at hand. Please check the Issues tab for more details.

## Reference
Some of the features, code and functions are directly referenced from https://nem-toolchain.github.io/. 

## Author
BramBear (Alvin Reyes)
