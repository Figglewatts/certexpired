.PHONY: tidy fmt build

build:
	go build -o bin/certexpired cmd/certexpired/main.go

tidy:
	go mod tidy

fmt:
	go fmt ./...