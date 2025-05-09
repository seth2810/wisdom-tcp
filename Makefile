.PHONY: lint test build_server run_server build_client run_client

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v2.1.0

lint: install-lint-deps
	golangci-lint run ./...

test:
	go test -count=1  ./...

build_server:
	go build -o ./bin/server ./cmd/server

run_server: build_server
	./bin/server

build_client:
	go build -o ./bin/client ./cmd/client

run_client: build_client
	./bin/client
