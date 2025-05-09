package server

type Config struct {
	ServerHost    string `envconfig:"SERVER_HOST"`
	ServerPort    uint16 `envconfig:"SERVER_PORT" default:"8080"`
	MinDifficulty uint8  `envconfig:"MIN_DIFFICULTY" default:"4"`
	MaxDifficulty uint8  `envconfig:"MAX_DIFFICULTY" default:"16"`
}
