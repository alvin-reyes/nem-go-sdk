package api

import (
	"fmt"
	"testing"
)

func TestGetAccount(t *testing.T) {
	a := NewAccountApi("http://alice2.nem.ninja:7890/")
	t.Log(a.AccountGet("NA3Y4G45LSCNMZZU6MDSAPSQQKPVQIW7JG4RP5PK"))
	fmt.Println(a.AccountGet("NA3Y4G45LSCNMZZU6MDSAPSQQKPVQIW7JG4RP5PK"))
}

func TestGenerateAccount(t *testing.T) {
	a := NewAccountApi("http://ec2-18-218-182-58.us-east-2.compute.amazonaws.com:7890/")
	//t.Log(a.AccountGenerate())
	b, _ := a.AccountGenerate()
	fmt.Println(b.Address)
	fmt.Println(b.PrivateKey)
	fmt.Println(b.PublicKey)
}
