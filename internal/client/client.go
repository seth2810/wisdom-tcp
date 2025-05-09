package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net"
	"strconv"
	"time"

	"github.com/seth2810/wisdom-tcp/internal/mtp"
	"github.com/seth2810/wisdom-tcp/internal/tcp"
)

type Client struct {
	cfg Config
}

func NewClient(cfg Config) *Client {
	return &Client{cfg: cfg}
}

func (c *Client) Request() error {
	addr := net.JoinHostPort(c.cfg.ServerHost, strconv.Itoa(int(c.cfg.ServerPort)))

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to establish connection %q: %w", addr, err)
	}

	//nolint:errcheck
	defer conn.Close()

	slog.Info("server connection established", slog.String("addr", addr))

	err = conn.SetDeadline(time.Now().Add(c.cfg.ConnectionDeadline))
	if err != nil {
		return fmt.Errorf("failed to set connection deadline: %w", err)
	}

	ch, err := c.readChallenge(conn)
	if err != nil {
		return fmt.Errorf("failed to read challenge: %w", err)
	}

	slog.Info("challenge received", slog.String("addr", addr), slog.Any("challenge", ch))

	pr, err := mtp.FindProof(mtp.DefaultConfig, ch)
	if err != nil {
		return fmt.Errorf("failed to find proof: %w", err)
	}

	slog.Info("proof found", slog.String("addr", addr), slog.Any("proof", pr))

	err = c.sendProof(conn, pr)
	if err != nil {
		return fmt.Errorf("failed to send proof: %w", err)
	}

	slog.Info("proof sent", slog.String("addr", addr))

	res, err := c.readResponse(conn)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	slog.Info("response received from server", slog.String("res", string(res)))

	return nil
}

func (c *Client) readChallenge(conn io.Reader) (*mtp.Challenge, error) {
	msg, err := tcp.ReadMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read challenge: %w", err)
	}

	var ch mtp.Challenge

	err = json.Unmarshal(msg, &ch)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal challenge: %w", err)
	}

	return &ch, nil
}

func (c *Client) sendProof(conn io.Writer, pr *mtp.Proof) error {
	msg, err := json.Marshal(pr)
	if err != nil {
		return fmt.Errorf("failed to marshal proof: %w", err)
	}

	err = tcp.WriteMessage(conn, msg)
	if err != nil {
		return fmt.Errorf("failed to write proof: %w", err)
	}

	return nil
}

func (c *Client) readResponse(conn io.Reader) ([]byte, error) {
	msg, err := tcp.ReadMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return msg, nil
}
