.DEFAULT_GOAL := all

all: lint test build

tool:
	@go install gotest.tools/gotestsum@latest
	@go install github.com/boumenot/gocover-cobertura@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/gobuild@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/golint@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/goupdate@latest

lint:
	@mkdir -p fixture/lint/empty_directory
	@golint --fix --skip fixture,tmp
	@golangci-lint run

test:
	@gotestsum --format standard-quiet -- ./...

coverage:
	@mkdir -p tmp
	@go test -coverprofile=tmp/coverage.txt -covermode count ./...
	@gocover-cobertura <tmp/coverage.txt >tmp/coverage.xml

update:
    # k8s.io/apimachinery: Cilium is not ready for 0.31.0
	@goupdate --downgrade k8s.io/apimachinery@v0.30.13

build:
	@go run cmd/gobuild/main.go --all

install:
	@go run cmd/gobuild/main.go --copy-to-bin gobuild
	@gobuild --all --copy-to-bin

chk:
	@gobuild --copy-to-bin gochk
	tmp/deploy.sh
