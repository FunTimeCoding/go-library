.DEFAULT_GOAL := all

all: test lint

lint:
	golangci-lint run

test:
	gotestsum --format standard-quiet -- ./...

dependencies:
	go get -u ./...

tools:
	go install gotest.tools/gotestsum@latest
