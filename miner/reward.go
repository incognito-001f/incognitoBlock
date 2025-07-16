package miner

func GetBlockReward(height int) uint64 {
	if height == 0 {
		return PremineRewardPerBlock
	}

	if height <= 50 {
		return PremineRewardPerBlock
	}

	return BlockReward
}
