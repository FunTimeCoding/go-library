.DEFAULT_GOAL := all

all: test lint

lint:
	golangci-lint run

test:
	gotestsum --format standard-quiet -- ./...

update:
	for ELEMENT in $$(go list -f "{{if not (or .Main .Indirect)}}{{.Path}}{{end}}" -m all); do echo $${ELEMENT}; go get $${ELEMENT}; done
	go mod tidy

tool:
	go install gotest.tools/gotestsum@latest

install:
	go build -o $$HOME/go/bin/bump cmd/bump/main.go

bump:
	go run cmd/bump/main.go
