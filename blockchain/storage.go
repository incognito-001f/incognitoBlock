package blockchain

import (
	"fmt"
	"os"
)

func SaveBlock(block *Block) {
	f, _ := os.OpenFile("blockchain.db", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	f.WriteString(fmt.Sprintf("Block %d - Hash: %s\n", block.Index, block.Hash))
	f.Close()
}
