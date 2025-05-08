package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/seth2810/wisdom-tcp/internal/quotes"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Listen(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen address %q: %w", address, err)
	}

	slog.Info("server started", slog.String("address", address))

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("failed to accept connection", slog.String("error", err.Error()))
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	//nolint:errcheck
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()

	slog.Info("client connected", slog.String("address", clientAddr))

	quote := quotes.GetRandomQuote()

	if _, err := conn.Write([]byte(quote)); err != nil {
		slog.Error("failed to send client response", slog.String("address", clientAddr), slog.String("error", err.Error()))
		return
	}

	slog.Info("response sent to client", slog.String("address", clientAddr))
}
