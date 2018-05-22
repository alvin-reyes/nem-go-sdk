package api

import (
	"testing"

	"github.com/nem-sdk-go/model"
)

func TestPrepareAnnounce(t *testing.T) {
	a := NewTransactionApi("http://alice2.nem.ninja:7890/")

	prepare := model.RequestPrepareAnnounce{}
	t.Log(a.TransactionPrepareAnnounce(prepare))

}

func TestAnnounce(t *testing.T) {
	a := NewTransactionApi("http://alice2.nem.ninja:7890/")

	requestAnnounce := model.RequestAnnounce{}
	t.Log(a.TransactionAnnounce(requestAnnounce))

}
