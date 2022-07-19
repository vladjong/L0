build:
	go build -v ./cmd/server
	./server

test:
	go test -v -race -timeout 30s ./...

migration:
	migrate -path migrations -database "postgres://localhost/l0?sslmode=disable" up

generate:
	go run publisher/publisher.go

nats:
	nats-streaming-server -cid prod -store file -dir store

clean:
	rm -rf server
	migrate -path migrations -database "postgres://localhost/l0?sslmode=disable" down