package main

import (
	"log/slog"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/seth2810/wisdom-tcp/internal/client"
)

func main() {
	var cfg client.Config

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if err := envconfig.Process("", &cfg); err != nil {
		slog.Error("failed to load env config", slog.String("err", err.Error()))
		os.Exit(1)
	}

	client := client.NewClient(cfg)

	if err := client.Request(); err != nil {
		slog.Error("failed to connect server", slog.String("err", err.Error()))
		os.Exit(1)
	}
}
