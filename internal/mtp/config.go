package mtp

var DefaultConfig = Config{
	TimeCost:    1,
	Parallelism: 4,
}

type Config struct {
	TimeCost    uint32 `json:"time_cost"`   // Number of iterations
	Parallelism uint8  `json:"parallelism"` // Number of parallel threads
}
