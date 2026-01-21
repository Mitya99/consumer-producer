GO_VERSION := 1.25

.PHONY: build test deps
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o kafka-consumer-producer .
	
deps:
	go mod download
