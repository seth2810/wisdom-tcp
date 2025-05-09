package main

import (
	"log/slog"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/seth2810/wisdom-tcp/internal/server"
)

func main() {
	var cfg server.Config

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if err := envconfig.Process("", &cfg); err != nil {
		slog.Error("failed to load env config", slog.String("err", err.Error()))
		os.Exit(1)
	}

	server := server.NewServer(cfg)
	if err := server.Listen(); err != nil {
		slog.Error("failed to start server", slog.String("err", err.Error()))
		os.Exit(1)
	}
}
