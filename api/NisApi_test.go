package api

import (
	"fmt"
	"testing"
)

func TestHeartBeat(t *testing.T) {
	a := NewNisApi("http://alice2.nem.ninja:7890/")
	t.Log(a.GetHeartBeat())
	fmt.Println(a.GetHeartBeat())
}

func TestStatus(t *testing.T) {
	a := NewNisApi("http://alice2.nem.ninja:7890/")
	t.Log(a.GetStatus())
	fmt.Println(a.GetStatus())
}
