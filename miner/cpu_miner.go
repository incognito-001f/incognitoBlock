package miner

import (
	"math/big"
)

func StartMiner(args []string) {
	nodeAddr := "localhost:8080"
	useGpu := false
	usePruned := false

	for _, arg := range args {
		if strings.HasPrefix(arg, "--node=") {
			nodeAddr = strings.TrimPrefix(arg, "--node=")
		}
		if arg == "--use-gpu" {
			useGpu = true
		}
		if arg == "--use-pruned" {
			usePruned = true
		}
	}

	if usePruned && !node.CanUsePruned() {
		fmt.Println("⚠️ El modo pruned aún no está habilitado")
		return
	}

	if useGpu {
		runGPUMiner(nodeAddr)
	} else {
		runCPUMiner(nodeAddr, usePruned)
	}
}
