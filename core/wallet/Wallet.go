package wallet

import (
	"encoding/base64"
	"encoding/json"

	"github.com/nem-sdk-go/core"
	"github.com/nem-sdk-go/core/keypair"
	"github.com/nem-sdk-go/core/wallet/account"
)

// NewWallet creates new default wallet.
func NewWallet() *Wallet {
	wlt := new(Wallet)
	wlt.Name = ""
	wlt.PrivateKey = ""
	wlt.Accounts = make(map[uint]*account.Account)

	return wlt
}

// Encode encodes wallet into base64 string.
func Encode(w *Wallet) (string, error) {
	var encoded string

	ser, err := json.Marshal(w)
	if err != nil {
		return encoded, err
	}
	encoded = base64.StdEncoding.EncodeToString(ser)

	return encoded, nil
}

// Decode decodes wallet form a base64 string.
func Decode(data string) (*Wallet, error) {
	wlt := NewWallet()

	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return wlt, err
	}
	err = json.Unmarshal(decoded, &wlt)
	return wlt, err
}

// Wallet represents wallet file.
type Wallet struct {
	PrivateKey string
	Name       string
	Accounts   map[uint]*account.Account
}

// AddAccount adds account into wallet.
func (wlt *Wallet) AddAccount(ch core.Chain, pair keypair.KeyPair, password string) error {
	acc := account.NewAccount(ch)
	err := acc.Encrypt(pair, password)
	if err != nil {
		return err
	}

	i := len(wlt.Accounts)
	wlt.Accounts[uint(i)] = acc

	return nil
}
