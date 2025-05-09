package client

import "time"

type Config struct {
	ServerHost         string        `envconfig:"SERVER_HOST"`
	ServerPort         uint16        `envconfig:"SERVER_PORT" default:"8080"`
	ConnectionDeadline time.Duration `envconfig:"CONNECTION_DEADLINE" default:"5s"`
}
