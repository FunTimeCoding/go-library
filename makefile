.DEFAULT_GOAL := all

all: test lint

lint:
	golangci-lint run

test:
	gotestsum --format standard-quiet -- ./...

install-dependencies:
	go install gotest.tools/gotestsum@latest
