.DEFAULT_GOAL := all

all: test lint build

tool:
	@go install gotest.tools/gotestsum@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/gobuild@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/golint@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/goupdate@latest

test:
	@gotestsum --format standard-quiet -- ./...

lint:
	@mkdir -p fixture/lint/empty_directory
	@golint --fix --skip fixture,tmp
	@golangci-lint run

update:
    # k8s.io/apimachinery: Cilium is not ready for 0.31.0
    # sigs.k8s.io/structured-merge-diff/v4: IgnoredFields error in structuredmerge.go
	@goupdate --downgrade k8s.io/apimachinery@v0.30.8 \
		--downgrade sigs.k8s.io/structured-merge-diff/v4@v4.4.1

build:
	@go run cmd/gobuild/main.go --all

clean:
	@rm --force --recursive tmp
	@mkdir tmp

install:
	@go run cmd/gobuild/main.go --copy-to-bin gobuild
	@gobuild --all --copy-to-bin

chk:
	@gobuild --copy-to-bin gochk
	tmp/deploy.sh
