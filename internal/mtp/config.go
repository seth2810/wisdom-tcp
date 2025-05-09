package mtp

var DefaultConfig = Config{
	TimeCost:      1,
	Parallelism:   4,
	MinDifficulty: 4,
	MaxDifficulty: 16,
}

type Config struct {
	TimeCost      uint32 `json:"time_cost"`      // Number of iterations
	Parallelism   uint8  `json:"parallelism"`    // Number of parallel threads
	MinDifficulty uint8  `json:"min_difficulty"` // Minimum difficulty level
	MaxDifficulty uint8  `json:"max_difficulty"` // Maximum difficulty level
}
