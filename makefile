.DEFAULT_GOAL := build

tool:
	@go install gotest.tools/gotestsum@latest
	@go install github.com/funtimecoding/go-library/cmd/goupdate@latest
	@go install github.com/funtimecoding/go-library/cmd/golint@latest
	@go install github.com/funtimecoding/go-library/cmd/gobuild@latest

test: tool
	@gotestsum --format standard-quiet -- ./...

lint: tool
	@golint
	@golangci-lint run

update: tool
	@for ELEMENT in $$(go list -f "{{if not (or .Main .Indirect)}}{{.Path}}{{end}}" -m all); do echo $${ELEMENT}; go get $${ELEMENT}; done
	@go mod tidy
	@goupdate

build: test lint
	@gobuild cmd/gobump/main.go
	@gobuild cmd/goclean/main.go
	@gobuild cmd/godebian/main.go
	@gobuild cmd/golint/main.go
	@gobuild cmd/goupdate/main.go

install: build
	@cp tmp/gobump $$HOME/go/bin/gobump
	@cp tmp/goclean $$HOME/go/bin/goclean
	@cp tmp/godebian $$HOME/go/bin/godebian
	@cp tmp/golint $$HOME/go/bin/golint
	@cp tmp/goupdate $$HOME/go/bin/goupdate
