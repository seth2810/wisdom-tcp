package client

import (
	"fmt"
	"log/slog"
	"net"
	"time"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to connect server: %w", err)
	}

	//nolint:errcheck
	defer conn.Close()

	err = conn.SetDeadline(time.Now().Add(30 * time.Second))
	if err != nil {
		return fmt.Errorf("failed to set connection deadline: %w", err)
	}

	slog.Info("connection established", slog.String("address", address))

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("failed to read server response: %w", err)
	}

	slog.Info("response received from server", slog.String("response", string(buf[:n])))

	return nil
}
