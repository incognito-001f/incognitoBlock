package blockchain

import (
	"fmt"
	"incognito-chain/utils"
	"time"
)

type Block struct {
	Index        int
	PreviousHash string
	Timestamp    int64
	Transactions []*Transaction
	Nonce        int
	Hash         string
	Difficulty   int
}

func NewBlock(index int, prevHash string, txs []*Transaction) *Block {
	return &Block{
		Index:        index,
		PreviousHash: prevHash,
		Timestamp:    time.Now().Unix(),
		Transactions: txs,
		Difficulty:   getDifficulty(index),
	}
}

func (b *Block) SetHash(nonce int) {
	b.Nonce = nonce
	data := fmt.Sprintf("%d%s%d%d", b.Index, b.PreviousHash, b.Timestamp, nonce)
	b.Hash = utils.CalculateHash(data)
}
