package node

import (
	"os"
)

// CanUsePruned verifica si se puede usar modo pruned
func CanUsePruned() bool {
	fi, err := os.Stat("blockchain.db")
	if err != nil {
		return false
	}
	return fi.Size() >= PruningThresholdSize
}
