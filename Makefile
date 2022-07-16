build:
	go build -v ./cmd/server

test:
	go test -v -race -timeout 30s ./...