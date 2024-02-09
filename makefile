.DEFAULT_GOAL := all

all: test lint

lint:
	@golangci-lint run

test:
	@gotestsum --format standard-quiet -- ./...

update:
	for ELEMENT in $$(go list -f "{{if not (or .Main .Indirect)}}{{.Path}}{{end}}" -m all); do echo $${ELEMENT}; go get $${ELEMENT}; done
	@go mod tidy

tool:
	go install gotest.tools/gotestsum@latest

install:
	go build -o $$HOME/go/bin/bump cmd/bump/main.go
	go build -o $$HOME/go/bin/lint cmd/lint/main.go
	go build -o $$HOME/go/bin/go-update cmd/update/main.go
	go build -o $$HOME/go/bin/clean-repository cmd/clean_repository/main.go

bump:
	@go run cmd/bump/main.go

build-debian-package:
	go build -o $$HOME/go/bin/debian-package cmd/debian_package/main.go
