package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/seth2810/wisdom-tcp/internal/client"
)

func main() {
	client := client.NewClient()
	address := net.JoinHostPort("", "8080")

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if err := client.Connect(address); err != nil {
		slog.Error("failed to connect server", slog.String("address", address), slog.String("error", err.Error()))
		os.Exit(1)
	}
}
