package vanity

import (
	"github.com/nem-sdk-go/core"
	"github.com/nem-sdk-go/core/keypair"
)

// StartSearch starts search process for vanity account address that satisfies a given selector.
func StartSearch(chain core.Chain, selector Selector, ch chan<- keypair.KeyPair) {
	for {
		pr := keypair.Gen()
		if !selector.Pass(pr.Address(chain)) {
			continue
		}
		ch <- pr
		return
	}
}
	