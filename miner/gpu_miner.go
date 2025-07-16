package miner

import "fmt"

type GPUMiner struct{}

func runGPUMiner(node string) {
	fmt.Printf("⛏️ [GPU] Conectando a %s\n", node)
	// Aquí va la lógica de conexión OpenCL/CUDA
}
