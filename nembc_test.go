package nembc

import (
	"fmt"
	"testing"
)

func TestNewNis(t *testing.T) {
	a := nembc.NewNemNis("http://alice2.nem.ninja:7890/",[]byte)
	t.Log(a.AccountGet("NA3Y4G45LSCNMZZU6MDSAPSQQKPVQIW7JG4RP5PK"))
	fmt.Println(a.AccountGet("NA3Y4G45LSCNMZZU6MDSAPSQQKPVQIW7JG4RP5PK"))
}
