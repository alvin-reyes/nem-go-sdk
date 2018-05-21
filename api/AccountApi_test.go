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
	a := NewAccountApi("http://alice2.nem.ninja:7890/")
	t.Log(a.AccountGenerate())
}
