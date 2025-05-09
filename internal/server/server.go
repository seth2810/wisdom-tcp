package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"
	"net"
	"strconv"
	"time"

	"github.com/seth2810/wisdom-tcp/internal/mtp"
	"github.com/seth2810/wisdom-tcp/internal/quotes"
	"github.com/seth2810/wisdom-tcp/internal/tcp"
)

type Server struct {
	cfg Config
}

func NewServer(cfg Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Listen() error {
	addr := net.JoinHostPort(s.cfg.ServerHost, strconv.Itoa(int(s.cfg.ServerPort)))

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen address %q: %w", addr, err)
	}

	slog.Info("server started", slog.String("addr", addr))

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("failed to accept connection", slog.String("err", err.Error()))
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	//nolint:errcheck
	defer conn.Close()

	addr := conn.RemoteAddr()
	clientIP, _, _ := net.SplitHostPort(addr.String())

	slog.Info("client connected", slog.String("addr", clientIP))

	difficulty := s.randomDifficulty(s.cfg.MinDifficulty, s.cfg.MaxDifficulty)

	nonce := mtp.GenerateNonce(difficulty)

	ch := &mtp.Challenge{
		Nonce:      nonce,
		Timestamp:  time.Now(),
		Difficulty: difficulty,
		SaltLength: 16,
		MemorySize: 8 * 1024 * 1024, //8MB
	}

	err := s.sendChallenge(conn, ch)
	if err != nil {
		slog.Error("error while sending challenge", slog.String("addr", clientIP), slog.String("err", err.Error()))
		return
	}

	slog.Info("challenge sent", slog.String("addr", clientIP))

	pr, err := s.readProof(conn)
	if err != nil {
		slog.Error("error while reading proof", slog.String("addr", clientIP), slog.String("err", err.Error()))
		return
	}

	slog.Info("proof received", slog.String("addr", clientIP))

	if ok := mtp.Verify(ch, pr); !ok {
		slog.Error("invalid proof received, rejecting connection", slog.String("addr", clientIP))
		return
	}

	slog.Info("proof verified", slog.String("addr", clientIP))

	res := quotes.GetRandomQuote()

	err = s.sendResponse(conn, res)
	if err != nil {
		slog.Error("error while sending response", slog.String("addr", clientIP), slog.String("err", err.Error()))
		return
	}

	slog.Info("response sent to client", slog.String("addr", clientIP))
}

func (s *Server) randomDifficulty(min, max uint8) uint8 {
	difficulty := rand.UintN(uint(max)-uint(min)) + uint(min)

	return uint8(difficulty)
}

func (s *Server) sendChallenge(conn io.Writer, ch *mtp.Challenge) error {
	msg, err := json.Marshal(ch)
	if err != nil {
		return fmt.Errorf("failed to marshal challenge: %w", err)
	}

	err = tcp.WriteMessage(conn, msg)
	if err != nil {
		return fmt.Errorf("failed to write challenge: %w", err)
	}

	return nil
}

func (s *Server) readProof(conn io.Reader) (*mtp.Proof, error) {
	msg, err := tcp.ReadMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read proof: %w", err)
	}

	var pr mtp.Proof

	if err := json.Unmarshal(msg, &pr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal proof: %w", err)
	}

	return &pr, nil
}

func (s *Server) sendResponse(conn io.Writer, data []byte) error {
	err := tcp.WriteMessage(conn, data)
	if err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}
