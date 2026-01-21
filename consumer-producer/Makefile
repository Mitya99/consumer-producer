GO_VERSION := 1.25

.PHONY: build test deps
build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -o kafka-consumer-producer .
	
deps:
	go mod download
