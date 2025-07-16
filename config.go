package main

const (
    TotalSupply               = 86_000_000 * TokenUnit
    TokenUnit                 = 1_000_000_000_000_000_000 // 1 token = 1e18 unidades
    PruningThresholdSize      = 6 * 1024 * 1024 * 1024     // 6 GB
    EnablePruningAfterHeight  = 10_000                     // Altura mínima para pruned
    BlockReward               = 6 * TokenUnit                // Recompensa por bloque
    PremineBlocks             = 50                          // Bloques iniciales
    PremineRewardPerBlock     = 120_000_000_000_000_000_000 // 120k tokens/bloque
)