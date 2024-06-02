build:
	go build -o bin/boi cmd/boi/main.go

fmt:
	go fmt ./...

test:
	go test ./...
