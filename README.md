# "Word of Wisdom" tcp server (protected from DDOS attacks with the Proof of Work).

## Task

Design and implement "Word of Wisdom" tcp server:

- TCP server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained.
- After Prof Of Work verification, server should send one of the quotes from "word of wisdom" book or any other collection of the quotes.
- Docker file should be provided both for the server and for the client that solves the POW challenge.

## Getting started

Requirements:

- Go 1.24+ installed (to run tests, start server or client without Docker)
- Docker installed (to run docker-compose)

```
# Run server and client by docker-compose
docker-compose up -d

# Run only server
make run_server

# Run only client
make run_client
```

## Resources

- [Word of Wisdom](https://en.wikipedia.org/wiki/Word_of_Wisdom)
- [Proof of work](https://en.wikipedia.org/wiki/Proof_of_work)

## Why Merkle Tree based PoW (MTP-like)?

I was looking for an algorithm that satisfies the following criteria:

- Easy to verify on the server
- Requires significant work on the client
- Uses both CPU and memory
- Dynamically adjustable difficulty
- Doesn't overly slow down real users
- Single-round interaction (no extra requests)

The Merkle Tree–based approach (in an MTP-lite style) is the only one that strikes a balanced middle ground, as it:

✅ Memory-bound work on client side (via Argon2 or similar):
Makes it expensive for attackers to parallelize or run on GPUs.

✅ Fast verification on server side:
Only a few hashes need to be computed to validate the Merkle root and challenge.

✅ Replay resistance:
Challenges include unique nonces and timestamps.

✅ Dynamic difficulty adjustment:
Server may increase difficulty under high load (e.g., requiring more hash matches), and decreases it under light load — adapting in real time.

✅ Scalable and asymmetric by design:
Clients spend seconds solving; server spends milliseconds verifying.

Trade-offs and Mitigations
Slightly more verification work on server side vs Hashcash:
Acceptable — 2–4 hash operations for tree verification is negligible vs the protection gained.

Challenge needs to include Merkle root + nonce + optional proof path:
Handled in protocol design with clear binary framing or JSON structure.

Higher RAM usage on client:
Configurable (e.g., 16–64 MB) and can be adjusted based on client capabilities.

Conclusion
The Merkle Tree–based PoW provides a modern, memory-bound, verifiable way to defend TCP servers from DDoS attacks.
It maintains the PoW principles of asymmetry, scalability, and difficulty tuning — while addressing many of the limitations found in older schemes like Hashcash.
For a high-performance, production-ready challenge-response server, this approach offers a balanced and future-proof solution.

## Project Structure

```
.
├── cmd/                    # Command-line applications
├── internal/              # Internal packages
│   ├── client/           # Client implementation
│   ├── mtp/              # Memory-hard Proof of Work implementation
│   ├── quotes/           # Quote collection
│   ├── server/           # Server implementation
│   └── tcp/              # TCP utilities
├── .docker/              # Docker-related files
└── docker-compose.yml    # Docker Compose configuration
```

## Implementation details

The client-server interaction follows a challenge-response protocol with the following steps:

### 1. Connection Establishment

- Client establishes TCP connection to server
- Server accepts connection and spawns a new goroutine to handle it
- Connection timeout is set (default: 5s for client)

### 2. Challenge Generation

Server generates a challenge containing:

```json
{
  "nonce": "16-byte random value",
  "timestamp": "current time",
  "difficulty": "random value between 4-16",
  "memory_size": "8MB",
  "salt_length": "16 bytes"
}
```

### 3. Proof of Work

Client performs the following steps:

1. Receives challenge from server
2. Generates memory-hard buffer using Argon2id
3. Builds Merkle tree from the buffer
4. Finds a nonce that satisfies the difficulty requirement
5. Sends proof containing:
   ```json
   {
     "root": "Merkle tree root hash",
     "nonce": "found nonce value"
   }
   ```

### 4. Verification

Server verifies the proof by:

1. Combining Merkle root and nonce
2. Calculating SHA-256 hash
3. Checking leading zero bits match difficulty
4. If valid, sends a random quote; if invalid, closes connection

### 5. Response

- On successful verification, server sends a random quote from the collection
- Connection is closed after quote delivery

### Message Format

All messages use a length-prefixed binary format:

1. 8-byte message length (uint64, big-endian)
2. Message payload (JSON for challenge/proof, plain text for quotes)

### Security Features

- Memory-hard computation prevents GPU-based attacks
- Dynamic difficulty adjustment (4-16 bits)
- Connection timeouts prevent resource exhaustion
- Unique nonces prevent replay attacks
- Merkle tree construction ensures proof integrity

## Security Considerations

- The server uses a memory-hard PoW algorithm to prevent GPU-based attacks
- Difficulty levels are configurable to adjust security requirements
- Connection timeouts are implemented to prevent resource exhaustion
- Random nonce generation ensures unique challenges

## License

This project is licensed under the terms of the license included in the repository.
