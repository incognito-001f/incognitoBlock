package wallet

import (
	"golang.org/x/crypto/ed25519"
)

type KeyPair struct {
	Public  ed25519.PublicKey
	Private ed25519.PrivateKey
}
