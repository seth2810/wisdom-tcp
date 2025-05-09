package tcp

import (
	"encoding/binary"
	"fmt"
	"io"
)

func ReadMessage(r io.Reader) ([]byte, error) {
	var ln uint64

	err := binary.Read(r, binary.BigEndian, &ln)
	if err != nil {
		return nil, fmt.Errorf("failed to read message size: %w", err)
	}

	buf := make([]byte, ln)

	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read message: %w", err)
	}

	return buf, nil
}

// WriteMessage writes a message to the connection.
func WriteMessage(w io.Writer, msg []byte) error {
	err := binary.Write(w, binary.BigEndian, uint64(len(msg)))
	if err != nil {
		return fmt.Errorf("failed to write message size: %w", err)
	}

	_, err = w.Write(msg)
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	return err
}
