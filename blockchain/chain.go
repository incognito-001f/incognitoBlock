package blockchain

import "sync"

var Chain []*Block
var mutex sync.Mutex

func LoadChain() *Blockchain {
	mutex.Lock()
	defer mutex.Unlock()

	if len(Chain) == 0 {
		genesis := NewBlock(0, "0", nil)
		genesis.SetHash(0)
		Chain = append(Chain, genesis)
		SaveBlock(genesis)
	}
	return &Blockchain{Chain: Chain}
}
