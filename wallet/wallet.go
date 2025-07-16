package wallet

import (
	"github.com/gtank/ristretto255"
)

type Wallet struct {
	PublicKey  *ristretto255.Element
	PrivateKey *ristretto255.Scalar
	Address    string
	Balance    uint64
}
