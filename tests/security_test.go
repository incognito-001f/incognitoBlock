package tests

import (
	"incognito-chain/blockchain"
	"testing"
)

func TestPremine(t *testing.T) {
	if blockchain.GetBlockReward(0) != PremineRewardPerBlock {
		t.FailNow()
	}
}

func TestBlockValidation(t *testing.T) {
	validBlock := blockchain.NewBlock(1, "validhash", nil)
	validBlock.SetHash(12345)
	if !blockchain.isBlockValid(validBlock) {
		t.FailNow()
	}

	modifiedBlock := blockchain.NewBlock(1, "invalidhash", nil)
	modifiedBlock.Hash = "falsohash123"
	if blockchain.isBlockValid(modifiedBlock) {
		t.FailNow()
	}
}
