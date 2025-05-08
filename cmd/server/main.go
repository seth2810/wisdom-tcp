package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/seth2810/wisdom-tcp/internal/server"
)

func main() {
	server := server.NewServer()
	address := net.JoinHostPort("", "8080")

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if err := server.Listen(address); err != nil {
		slog.Error("failed to start server", slog.String("address", address), slog.String("error", err.Error()))
		os.Exit(1)
	}
}
