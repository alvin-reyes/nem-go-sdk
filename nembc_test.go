package nembc

import (
  "fmt"
  "testing"
)

func TestNewNis(t *testing.T) {
  a := NewNisInstance("http://alice2.nem.ninja:7890/", Testnet)
  fmt.Println("===========Get account information============" )
  fmt.Println(a.AccountApi.AccountGet("NA3Y4G45LSCNMZZU6MDSAPSQQKPVQIW7JG4RP5PK"))
  fmt.Println("===========Get namespace and mosaic information===========")
  fmt.Println(a.NamespaceMosaicsApi.GetNamespaceRootPage("rootpage"))
  fmt.Println("===========Blockchain Api endpoint===========")
  fmt.Println(a.BlockchainApi.GetChainHeight())
  fmt.Println("===========NisApi Heart Beat===========")
  fmt.Println(a.NisApi.GetHeartBeat())
  fmt.Println("===========Node Information===========")
  fmt.Println(a.NodeApi.GetNodeInfo())
  // Transaction Api
  // fmt.Println(a.TransactionApi.TransactionPrepareAnnounce())
}
